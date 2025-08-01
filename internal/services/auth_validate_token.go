package services

import (
	"context"
	"errors"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/security"
)

// Convertir tus constantes de string a errores
var (
	ErrInvalidToken    = errors.New("invalid_token")
	ErrSessionNotFound = errors.New("session_not_found")
	ErrSessionMismatch = errors.New("session_mismatch")
	ErrSessionInactive = errors.New("session_inactive")
)

// ValidateToken sirve para validar un access token y su sesión asociada.
func (s *AuthService) ValidateToken(ctx context.Context, input dto.TokenValidationRequestDTO) (*dto.TokenValidationResponseDTO, error) {
	// 1 Obtener token por hash
	tokModel, err := s.tokenRepo.FindByHash(ctx, input.TokenHash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2 Estado del token debe ser activo
	if tokModel.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}

	// 3 Obtener sesión asociada
	sessModel, err := s.sessionRepo.FindBySessionID(ctx, tokModel.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	// 4 Coincidencia de session ID
	if sessModel.SessionID != input.SessionID {
		return nil, ErrSessionMismatch
	}

	// 5 Verificar estado de sesión
	if sessModel.Status != models.SessionStatusActive || !sessModel.IsActive {
		return nil, ErrSessionInactive
	}

	// 6 Validar JWE y extraer claims
	val := security.ValidateToken(input.TokenHash)
	details := dto.TokenValidationDetailsResponseDTO{}
	switch val.Code {
	case 0:
		details.Valid = true
		details.Message = "Token válido"
		claims := val.Claims
		details.Subject = claims.Subject
		details.IssuedAt = claims.IssuedAt.Time().Format(time.RFC3339)
		details.ExpiresAt = claims.Expiry.Time().Format(time.RFC3339)
		details.ExpiresIn = int64(time.Until(claims.Expiry.Time()).Seconds())
	case 2:
		details.Valid = false
		details.Message = "Token expirado"
	default:
		details.Valid = false
		details.Message = val.Message
	}

	// 7 Construir DTO de respuesta
	return &dto.TokenValidationResponseDTO{
		UserID:      tokModel.UserID,
		TokenID:     tokModel.TokenID,
		SessionID:   tokModel.SessionID,
		Status:      tokModel.Status,
		TokenType:   tokModel.TokenType,
		TokenDetail: details,
	}, nil
}
