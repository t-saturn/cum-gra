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

// Introspect valida el token contra DB y JWS, y devuelve metadata del token.
func (s *AuthService) Introspect(ctx context.Context, rawToken string, q dto.IntrospectQueryDTO) (*dto.TokenIntrospectResponseDTO, error) {
	// 0. Validación mínima
	if rawToken == "" || q.SessionID == "" {
		return nil, ErrInvalidToken
	}

	// 1. Lookup por hash del token crudo
	calcHash := security.HashTokenHex(rawToken)
	tokModel, err := s.tokenRepo.FindByHash(ctx, calcHash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2. Debe estar activo
	if tokModel.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}

	// Expiración por DB (idempotente)
	now := time.Now().UTC()
	if now.After(tokModel.ExpiresAt) {
		_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, now)
		return nil, security.ErrTokenExpired
	}

	// 3. Cargar sesión asociada
	sessModel, err := s.sessionRepo.FindBySessionID(ctx, tokModel.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	// 4. Coincidencia de session_id
	if sessModel.SessionID != q.SessionID {
		return nil, ErrSessionMismatch
	}

	// 5. Verificar estado de sesión
	if sessModel.Status != models.SessionStatusActive || !sessModel.IsActive {
		return nil, ErrSessionInactive
	}

	// 6. Defensa anti-swap: el hash calculado debe coincidir con el almacenado
	if !security.CompareHash(calcHash, tokModel.TokenHash) {
		return nil, ErrInvalidToken
	}

	// 7. Verificar JWS (RS256) y extraer claims
	claims, vErr := security.VerifyTokenRS256(rawToken)
	if vErr != nil {
		if errors.Is(vErr, security.ErrTokenExpired) {
			if tokModel.Status == models.TokenStatusActive {
				_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, time.Now().UTC())
			}
			return nil, security.ErrTokenExpired
		}
		return nil, security.ErrTokenInvalid
	}

	// 8. Construir detalles
	details := dto.IntrospectDetailsResponseDTO{
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

	// 9. Respuesta
	return &dto.TokenIntrospectResponseDTO{
		UserID:      tokModel.UserID,
		TokenID:     tokModel.TokenID, // o tokModel.ID.Hex() si usas ObjectID
		SessionID:   tokModel.SessionID,
		Status:      tokModel.Status,
		TokenType:   tokModel.TokenType,
		TokenDetail: details,
	}, nil
}
