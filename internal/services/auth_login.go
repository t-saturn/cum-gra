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
)

// Login realiza el flujo completo de autenticación
func (s *AuthService) Login(ctx context.Context, input dto.AuthLoginRequestDTO) (*dto.AuthLoginResponseDTO, error) {
	now := utils.NowUTC()

	// 1) Intento: usuario no encontrado / inactivo
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

	// 2) Intento: contraseña inválida
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		if _, recErr := s.InsertAttempt(ctx, input, models.AuthStatusFailed, models.AuthStatusInvalidPass, user.ID.String()); recErr != nil {
			logger.Log.Errorf("Error guardando AuthAttempt (bad password): %v", recErr)
		}
		return nil, ErrInvalidCredentials
	}

	// 3) Intento exitoso
	authAttemptID, err := s.InsertAttempt(ctx, input, models.AuthStatusSuccess, models.AuthStatusSuccess, user.ID.String())
	if err != nil {
		logger.Log.Errorf("Error guardando AuthAttempt (success): %v", err)
	}

	// 4) Crear sesión
	_, sessionDTO, err := s.InsertSession(ctx, input, user.ID.String(), authAttemptID, now)
	if err != nil {
		return nil, err
	}

	// 5) Generar y persistir tokens
	expMin, _ := strconv.Atoi(config.GetConfig().Server.JWTExpMinutes)
	durAccess := time.Duration(expMin) * time.Minute
	accessJWT, _ := security.GenerateAccessToken(user.ID.String())
	accessMongoID, _ := s.InsertToken(ctx, input, user.ID.String(), sessionDTO.SessionID, now, models.TokenTypeAccess, durAccess, nil)
	accessDetail := dto.TokenDetailDTO{
		TokenID:   accessMongoID.Hex(),
		Token:     accessJWT,
		TokenType: models.TokenTypeAccess,
		ExpiresAt: now.Add(durAccess),
	}

	durRefresh := 7 * 24 * time.Hour
	refreshJWT, _ := security.GenerateRefreshToken(user.ID.String())
	refreshMongoID, _ := s.InsertToken(ctx, input, user.ID.String(), sessionDTO.SessionID, now, models.TokenTypeRefresh, durRefresh, &accessMongoID)
	refreshDetail := dto.TokenDetailDTO{
		TokenID:   refreshMongoID.Hex(),
		Token:     refreshJWT,
		TokenType: models.TokenTypeRefresh,
		ExpiresAt: now.Add(durRefresh),
	}

	// 6) Devolver respuesta
	return &dto.AuthLoginResponseDTO{
		UserID:    user.ID.String(),
		Session:   sessionDTO,
		Tokens:    dto.TokensDTO{AccessToken: accessDetail, RefreshToken: refreshDetail},
		AttemptID: authAttemptID.Hex(),
	}, nil
}
