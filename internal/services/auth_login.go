package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrInactiveAccount    = errors.New("cuenta inactiva")
)

// AuthService gestiona la lógica de autenticación.
type AuthService struct {
	userRepo        *repositories.UserRepository
	authAttemptRepo *repositories.AuthAttemptRepository
	sessionRepo     *repositories.SessionRepository
	tokenRepo       *repositories.TokenRepository
}

// NewAuthService construye un AuthService con Postgres y Mongo ya conectados.
func NewAuthService(pgDB *gorm.DB, mongoDB *mongo.Database) *AuthService {
	return &AuthService{
		userRepo:        repositories.NewUserRepository(pgDB),
		authAttemptRepo: repositories.NewAuthAttemptRepository(mongoDB),
		sessionRepo:     repositories.NewSessionRepository(mongoDB),
		tokenRepo:       repositories.NewTokenRepository(mongoDB),
	}
}

// Login realiza el flujo completo de autenticación
func (s *AuthService) Login(ctx context.Context, input dto.AuthLoginRequestDTO) (*dto.AuthLoginResponseDTO, error) {
	now := utils.NowUTC()

	// 1 Intento: usuario no encontrado / inactivo
	user, err := s.userRepo.FindActiveByEmailOrDNI(ctx, &input.Email, nil)
	if err != nil {
		reason := models.AuthStatusInvalidUser // "invalid_credentials"
		if errors.Is(err, repositories.ErrUserDeleted) || errors.Is(err, repositories.ErrUserDisabled) {
			reason = "account_inactive"
		}
		// grabamos siempre el intento
		if _, recErr := s.InsertAttempt(ctx, input, models.AuthStatusFailed, reason, ""); recErr != nil {
			logger.Log.Errorf("Error guardando AuthAttempt (lookup): %v", recErr)
		}
		// devolvemos el error al usuario
		if reason == "account_inactive" {
			return nil, ErrInactiveAccount
		}
		return nil, ErrInvalidCredentials
	}

	// 2 Intento: contraseña inválida
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		if _, recErr := s.InsertAttempt(ctx, input, models.AuthStatusFailed, models.AuthStatusInvalidPass, user.ID.String()); recErr != nil {
			logger.Log.Errorf("Error guardando AuthAttempt (bad password): %v", recErr)
		}
		return nil, ErrInvalidCredentials
	}

	// 3 Intento exitoso
	authAttemptID, err := s.InsertAttempt(ctx, input, models.AuthStatusSuccess, models.AuthStatusSuccess, user.ID.String())
	if err != nil {
		logger.Log.Errorf("Error guardando AuthAttempt (success): %v", err)
	}

	// 4 Crear sesión
	_, sessionDTO, err := s.InsertSession(ctx, input, user.ID.String(), authAttemptID, now)
	if err != nil {
		return nil, err
	}

	// 5 Generar y persistir tokens
	expMin, _ := strconv.Atoi(config.GetConfig().Server.JWTExpMinutes)
	durAccess := time.Duration(expMin) * time.Minute

	accessID, accessJWT, err := s.InsertToken(ctx, user.ID.String(), sessionDTO.SessionID, input.DeviceInfo, models.TokenTypeAccess, durAccess, nil)
	if err != nil {
		return nil, err
	}
	accessDetail := dto.TokenLoginDTO{
		TokenID:   accessID.Hex(),
		Token:     accessJWT,
		TokenType: models.TokenTypeAccess,
		ExpiresAt: now.Add(durAccess),
	}

	durRefresh := 7 * 24 * time.Hour
	refreshID, refreshJWT, err := s.InsertToken(ctx, user.ID.String(), sessionDTO.SessionID, input.DeviceInfo, models.TokenTypeRefresh, durRefresh, nil)
	if err != nil {
		return nil, err
	}
	refreshDetail := dto.TokenLoginDTO{
		TokenID:   refreshID.Hex(),
		Token:     refreshJWT,
		TokenType: models.TokenTypeRefresh,
		ExpiresAt: now.Add(durRefresh),
	}

	// 6 Vincular tokens (access y refresh)
	if err := s.tokenRepo.SetPairedTokenID(ctx, refreshID, accessID); err != nil {
		logger.Log.Errorf("Error al vincular refresh_token con access_token: %v", err)
	}
	if err := s.tokenRepo.SetPairedTokenID(ctx, accessID, refreshID); err != nil {
		logger.Log.Errorf("Error al vincular access_token con refresh_token: %v", err)
	}

	// 7 Asociar tokens a la sesión
	if err := s.sessionRepo.AddTokenToSession(ctx, sessionDTO.SessionID, accessID); err != nil {
		logger.Log.Errorf("Error al asociar access_token a la sesión: %v", err)
	}
	if err := s.sessionRepo.AddTokenToSession(ctx, sessionDTO.SessionID, refreshID); err != nil {
		logger.Log.Errorf("Error al asociar refresh_token a la sesión: %v", err)
	}

	// 8 Devolver respuesta
	return &dto.AuthLoginResponseDTO{
		UserID:  user.ID.String(),
		Session: sessionDTO,
		Tokens: dto.TokensLoginDTO{
			AccessToken:  accessDetail,
			RefreshToken: refreshDetail,
		},
		AttemptID: authAttemptID.Hex(),
	}, nil
}
