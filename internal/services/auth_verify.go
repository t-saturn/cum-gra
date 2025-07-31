package services

import (
	"context"
	"errors"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repositories"
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
	userRepo          *repositories.UserRepository
	authAttemptRepo   *repositories.AuthAttemptRepository
	verifyAttemptRepo *repositories.VerifyAttemptRepository
	sessionRepo       *repositories.SessionRepository
	tokenRepo         *repositories.TokenRepository
}

// NewAuthService construye un AuthService con Postgres y Mongo ya conectados.
func NewAuthService(pgDB *gorm.DB, mongoDB *mongo.Database) *AuthService {
	return &AuthService{
		userRepo:          repositories.NewUserRepository(pgDB),
		authAttemptRepo:   repositories.NewAuthAttemptRepository(mongoDB),
		verifyAttemptRepo: repositories.NewVerifyAttemptRepository(mongoDB),
		sessionRepo:       repositories.NewSessionRepository(mongoDB),
		tokenRepo:         repositories.NewTokenRepository(mongoDB),
	}
}

// VerifyCredentials verifica email/DNI + contraseña y retorna el wrapper genérico.
func (s *AuthService) VerifyCredentials(ctx context.Context, input dto.AuthVerifyRequestDTO) (*dto.ResponseDTO[dto.AuthVerifyResponseDTO], error) {
	start := time.Now()

	// 1 Cargar usuario
	userData, err := s.userRepo.FindActiveByEmailOrDNI(ctx, input.Email, input.DNI)
	if err != nil {
		reason := ""
		switch {
		case errors.Is(err, repositories.ErrUserDeleted), errors.Is(err, repositories.ErrUserDisabled):
			reason = "account_inactive"
			s.InsertVerify(ctx, input, models.AuthStatusFailed, reason, "", 0)
			return nil, ErrInactiveAccount
		case errors.Is(err, gorm.ErrRecordNotFound), errors.Is(err, repositories.ErrUserNotFound):
			reason = "user_not_found"
			s.InsertVerify(ctx, input, models.AuthStatusFailed, reason, "", 0)
			return nil, ErrInvalidCredentials
		default:
			return nil, err
		}
	}

	// 2 Verificar la contraseña
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, userData.PasswordHash) {
		s.InsertVerify(ctx, input, models.AuthStatusFailed, "invalid_password", "", 0)
		return nil, ErrInvalidCredentials
	}
	elapsed := time.Since(start).Milliseconds()

	// 3 Registrar intento exitoso
	// (La función InsertVerify debería insertar y devolver el ObjectID si lo necesitas como AttemptID)
	objID, err := s.InsertVerify(ctx, input, models.AuthStatusSuccess, "correct", userData.ID.String(), elapsed)
	if err != nil {
		return nil, err
	}

	// 4 Construir el DTO específico
	now := utils.NowUTC()
	data := dto.AuthVerifyResponseDTO{
		AttemptID:   objID.Hex(),
		UserID:      userData.ID.String(),
		Status:      models.AuthStatusSuccess,
		ValidatedAt: now,
		ValidationResponse: dto.ValidationResponseDTO{
			UserID:          userData.ID.String(),
			ServiceResponse: models.AuthStatusSuccess,
			ValidatedBy:     models.AuthMethodCredentials,
			ValidationTime:  elapsed,
		},
	}

	// 5 Envolverlo en el ResponseDTO genérico
	resp := &dto.ResponseDTO[dto.AuthVerifyResponseDTO]{
		Success: true,
		Message: "Credenciales válidas",
		Data:    data,
	}
	return resp, nil
}
