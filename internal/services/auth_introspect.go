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
	ErrInvalidToken          = errors.New("invalid_token")
	ErrSessionMismatch       = errors.New("session_mismatch")
	ErrMissingFields         = errors.New("missing_fields")
	ErrSessionNotFound       = errors.New("session_not_found")
	ErrSessionInactive       = errors.New("session_inactive")
	ErrTokenNotFound         = errors.New("token_not_found")
	ErrTokenSessionMismatch  = errors.New("token_session_mismatch")
	ErrTokenTypeMismatch     = errors.New("token_type_mismatch")
	ErrTokenRevoked          = errors.New("token_revoked")
	ErrMalformedAccessToken  = errors.New("malformed_access_token")
	ErrMalformedRefreshToken = errors.New("malformed_refresh_token")
)

// Introspect valida el token contra DB y JWS, y devuelve metadata del token.
func (s *AuthService) IntrospectSessionTokens(ctx context.Context, sessionID, rawAccess, rawRefresh string) (*dto.IntrospectResponseDTO, error) {
	if sessionID == "" || rawAccess == "" || rawRefresh == "" {
		return nil, ErrMissingFields
	}

	// 1. Leer sesión
	sess, err := s.sessionRepo.FindBySessionID(ctx, sessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}
	if sess == nil {
		return nil, ErrSessionNotFound
	}

	if sess.Status != models.SessionStatusActive || !sess.IsActive || time.Now().UTC().After(sess.ExpiresAt) {
		return nil, ErrSessionInactive
	}

	// 2. Construir vistas (usa SOLO VerifyTokenRS256 internamente)
	accView, err := s.buildTokenView(ctx, rawAccess, models.TokenTypeAccess, sessionID)
	if err != nil {
		return nil, err
	}

	refView, err := s.buildTokenView(ctx, rawRefresh, models.TokenTypeRefresh, sessionID)
	if err != nil {
		return nil, err
	}

	// 3. Respuesta
	resp := &dto.IntrospectResponseDTO{
		UserID:           accView.TokenDetail.Subject, // o desde tu modelo (tok.UserID) si prefieres
		SessionID:        sessionID,
		Status:           sess.Status,
		SessionExpiresAt: sess.ExpiresAt.Format(time.RFC3339),
		Tokens: dto.TokensDTO{
			AccessToken:  accView,
			RefreshToken: refView,
		},
	}

	return resp, nil
}

func (s *AuthService) buildTokenView(ctx context.Context, rawToken string, expectedType string, sessionID string) (dto.TokenViewDTO, error) {
	// 1. Lookup por hash + tipo (evita swap entre access/refresh)
	hash := security.HashTokenHex(rawToken)

	tok, err := s.tokenRepo.FindByHashAndType(ctx, hash, expectedType)
	if err != nil {
		return dto.TokenViewDTO{}, ErrTokenNotFound
	}
	if tok == nil {
		return dto.TokenViewDTO{}, ErrTokenNotFound
	}

	// 2. Debe pertenecer a la sesión pedida
	if tok.SessionID != sessionID {
		return dto.TokenViewDTO{}, ErrTokenSessionMismatch
	}

	// 3. Verificar firma/claims con go-jose (única fuente)
	claims, vErr := security.VerifyTokenRS256(rawToken)

	now := time.Now().UTC()
	status := tok.Status

	// 4. Interpretación del error de verificación
	switch {
	case vErr == nil:
	case errors.Is(vErr, security.ErrTokenExpired):
		if status == models.TokenStatusActive {
			if err := s.tokenRepo.MarkExpired(ctx, tok.ID, now); err != nil {
			} else {
			}
			status = models.TokenStatusExpired
		}
	default:
		if expectedType == models.TokenTypeAccess {
			return dto.TokenViewDTO{}, ErrMalformedAccessToken
		}
		return dto.TokenViewDTO{}, ErrMalformedRefreshToken
	}

	// 5. Si DB dice revoked, eso sí es error (401)
	if status == models.TokenStatusRevoked {
		return dto.TokenViewDTO{}, ErrTokenRevoked
	}

	// 6. Construir detalles desde claims
	var issuedAt, expiresAt string
	var expiresIn int64

	if claims.IssuedAt != nil {
		issuedAt = claims.IssuedAt.Time().Format(time.RFC3339)
	}
	if claims.Expiry != nil {
		exp := claims.Expiry.Time()
		expiresAt = exp.Format(time.RFC3339)
		expiresIn = int64(time.Until(exp).Seconds())
	}

	detail := dto.IntrospectDetailsResponseDTO{
		Valid:     true,
		Message:   map[bool]string{true: "Token válido (expirado)", false: "Token válido"}[errors.Is(vErr, security.ErrTokenExpired)],
		Subject:   claims.Subject,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
		ExpiresIn: expiresIn,
	}

	result := dto.TokenViewDTO{
		TokenID:     tok.TokenID,
		Status:      status,
		TokenType:   tok.TokenType,
		TokenDetail: detail,
	}

	return result, nil
}
