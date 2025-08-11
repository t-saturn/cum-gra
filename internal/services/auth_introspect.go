package services

import (
	"context"
	"errors"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/security"
)

// Errores sentinela del service (sin cambios)
var (
	ErrInvalidToken    = errors.New("invalid_token")
	ErrSessionNotFound = errors.New("session_not_found")
	ErrSessionMismatch = errors.New("session_mismatch")
	ErrSessionInactive = errors.New("session_inactive")
)

func (s *AuthService) Introspect(ctx context.Context, input dto.IntrospectRequestDTO) (*dto.TokenIntrospectResponseDTO, error) {
	// Validación mínima de entrada
	if input.Token == "" || input.SessionID == "" {
		return nil, ErrInvalidToken
	}

	// 1. Calcular hash del token crudo y buscar en DB (lookup rápido)
	calcHash := security.HashTokenHex(input.Token)
	tokModel, err := s.tokenRepo.FindByHash(ctx, calcHash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2. Estado del token debe ser activo
	if tokModel.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}

	// 2.5) Expiración por DB
	now := time.Now().UTC()
	if now.After(tokModel.ExpiresAt) {
		_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, now)
		return nil, security.ErrTokenExpired
	}

	// 3. Obtener sesión asociada
	sessModel, err := s.sessionRepo.FindBySessionID(ctx, tokModel.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	// 4. Coincidencia de session ID
	if sessModel.SessionID != input.SessionID {
		return nil, ErrSessionMismatch
	}

	// 5. Verificar estado de sesión
	if sessModel.Status != models.SessionStatusActive || !sessModel.IsActive {
		return nil, ErrSessionInactive
	}

	// 6. Confirmar que el hash recalculado coincida con el almacenado (defensa anti-swap)
	if !security.CompareHash(calcHash, tokModel.TokenHash) {
		return nil, ErrInvalidToken
	}

	// 7. Verificar JWS (RS256) y extraer claims sobre el token crudo
	claims, vErr := security.VerifyTokenRS256(input.Token)
	if vErr != nil {
		if errors.Is(vErr, security.ErrTokenExpired) {
			if tokModel.Status == models.TokenStatusActive {
				_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, time.Now().UTC())
			}
			return nil, security.ErrTokenExpired
		}
		return nil, security.ErrTokenInvalid
	}

	details := dto.IntrospectDetailsResponseDTO{}
	// 7) Verificar JWS (RS256) y extraer claims sobre el token crudo
	claims, vErr = security.VerifyTokenRS256(input.Token)
	if vErr != nil {
		if errors.Is(vErr, security.ErrTokenExpired) {
			if tokModel.Status == models.TokenStatusActive {
				_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, time.Now().UTC())
			}
			return nil, security.ErrTokenExpired
		}
		return nil, security.ErrTokenInvalid
	}

	// 7.1) Si llegamos aquí, el token es válido (vErr == nil)
	details = dto.IntrospectDetailsResponseDTO{
		Valid:   true,
		Message: "Token válido",
		Subject: claims.Subject,
	}
	if claims.IssuedAt != nil {
		details.IssuedAt = claims.IssuedAt.Time().Format(time.RFC3339)
	}
	if claims.Expiry != nil {
		expTime := claims.Expiry.Time()
		details.ExpiresAt = expTime.Format(time.RFC3339)
		sec := int64(time.Until(expTime).Seconds())
		if sec < 0 {
			sec = 0
		}
		details.ExpiresIn = sec
	}

	// 8) Respuesta
	return &dto.TokenIntrospectResponseDTO{
		UserID:      tokModel.UserID,
		TokenID:     tokModel.TokenID,
		SessionID:   tokModel.SessionID,
		Status:      tokModel.Status,
		TokenType:   tokModel.TokenType,
		TokenDetail: details,
	}, nil
}
