package services

import (
	"context"
	"strconv"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// Login realiza el flujo completo de autenticación:
// 1) verifica credenciales,
// 2) registra AuthAttempt,
// 3) crea Session,
// 4) genera y persiste Access y Refresh tokens,
// 5) devuelve el DTO con session, tokens y attempt_id.
func (s *AuthService) Login(ctx context.Context, input dto.AuthLoginRequestDTO) (*dto.AuthLoginResponseDTO, error) {
	now := utils.NowUTC()

	// 1) Verificar usuario y contraseña
	user, err := s.userRepo.FindActiveByEmailOrDNI(ctx, &input.Email, nil)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// 2) Registrar intento de autenticación
	authAttemptID, err := s.InsertAttempt(ctx, dto.AuthVerifyRequestDTO{Email: &input.Email, ApplicationID: input.ApplicationID, DeviceInfo: input.DeviceInfo}, models.AuthStatusSuccess, user.ID.String())
	if err != nil {
		logger.Log.Errorf("Error guardando AuthAttempt: %v", err)
	}

	// 3) Crear sesión
	sessionID, sessionDTO, err := s.InsertSession(ctx, input, user.ID.String(), authAttemptID, now)
	if err != nil {
		return nil, err
	}

	// 4) Generar y persistir tokens
	//    – Access
	expMin, _ := strconv.Atoi(config.GetConfig().Server.JWTExpMinutes)
	durAccess := time.Duration(expMin) * time.Minute
	accessJWT, err := security.GenerateAccessToken(user.ID.String())
	if err != nil {
		return nil, err
	}
	accessMongoID, err := s.InsertToken(ctx, input, user.ID.String(), sessionID.String(), now, models.TokenTypeAccess, durAccess, nil)
	if err != nil {
		return nil, err
	}
	accessDetail := dto.TokenDetailDTO{
		TokenID:   accessMongoID.Hex(),
		Token:     accessJWT,
		TokenType: models.TokenTypeAccess,
		ExpiresAt: now.Add(durAccess),
	}

	//    – Refresh
	durRefresh := 7 * 24 * time.Hour
	refreshJWT, err := security.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}
	refreshMongoID, err := s.InsertToken(ctx, input, user.ID.String(), sessionID.String(), now, models.TokenTypeRefresh, durRefresh, &accessMongoID)
	if err != nil {
		return nil, err
	}
	refreshDetail := dto.TokenDetailDTO{
		TokenID:   refreshMongoID.Hex(),
		Token:     refreshJWT,
		TokenType: models.TokenTypeRefresh,
		ExpiresAt: now.Add(durRefresh),
	}

	// 5) Construir y devolver DTO de respuesta
	return &dto.AuthLoginResponseDTO{
		UserID:    user.ID.String(),
		Session:   sessionDTO,
		Tokens:    dto.TokensDTO{AccessToken: accessDetail, RefreshToken: refreshDetail},
		AttemptID: authAttemptID.Hex(),
	}, nil
}
