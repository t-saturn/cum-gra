package services

import (
	"context"
	"errors"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/logger"
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

	logger.Log.Debugf("[IntrospectSessionTokens] start | sessionID=%s", sessionID)

	if sessionID == "" || rawAccess == "" || rawRefresh == "" {
		logger.Log.Errorf("[IntrospectSessionTokens] missing fields | sessionID=%s rawAccess.len=%d rawRefresh.len=%d",
			sessionID, len(rawAccess), len(rawRefresh))
		return nil, ErrMissingFields
	}

	// 1. Leer sesión
	sess, err := s.sessionRepo.FindBySessionID(ctx, sessionID)
	if err != nil {
		logger.Log.Errorf("[IntrospectSessionTokens] error buscando sesión en repo: %v", err)
		return nil, ErrSessionNotFound
	}
	if sess == nil {
		logger.Log.Debugf("[IntrospectSessionTokens] sesión no encontrada | sessionID=%s", sessionID)
		return nil, ErrSessionNotFound
	}

	logger.Log.Debugf("[IntrospectSessionTokens] sesión encontrada | sessionID=%s status=%s isActive=%v expiresAt=%s",
		sess.SessionID, sess.Status, sess.IsActive, sess.ExpiresAt.Format(time.RFC3339))

	if sess.Status != models.SessionStatusActive || !sess.IsActive || time.Now().UTC().After(sess.ExpiresAt) {
		logger.Log.Debugf("[IntrospectSessionTokens] sesión inactiva o expirada | status=%s isActive=%v now=%s expiresAt=%s",
			sess.Status, sess.IsActive, time.Now().UTC().Format(time.RFC3339), sess.ExpiresAt.Format(time.RFC3339))
		return nil, ErrSessionInactive
	}

	// 2. Construir vistas (usa SOLO VerifyTokenRS256 internamente)
	accView, err := s.buildTokenView(ctx, rawAccess, models.TokenTypeAccess, sessionID)
	if err != nil {
		logger.Log.Errorf("[IntrospectSessionTokens] error construyendo vista access_token: %v", err)
		return nil, err
	}
	logger.Log.Debugf("[IntrospectSessionTokens] access_token view OK | tokenID=%s status=%s subject=%s",
		accView.TokenID, accView.Status, accView.TokenDetail.Subject)

	refView, err := s.buildTokenView(ctx, rawRefresh, models.TokenTypeRefresh, sessionID)
	if err != nil {
		logger.Log.Errorf("[IntrospectSessionTokens] error construyendo vista refresh_token: %v", err)
		return nil, err
	}
	logger.Log.Debugf("[IntrospectSessionTokens] refresh_token view OK | tokenID=%s status=%s subject=%s",
		refView.TokenID, refView.Status, refView.TokenDetail.Subject)

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

	logger.Log.Debugf("[IntrospectSessionTokens] success | userID=%s sessionID=%s status=%s sessionExpiresAt=%s",
		resp.UserID, resp.SessionID, resp.Status, resp.SessionExpiresAt)

	return resp, nil
}

func (s *AuthService) buildTokenView(ctx context.Context, rawToken string, expectedType string, sessionID string) (dto.TokenViewDTO, error) {

	logger.Log.Debugf("[buildTokenView] start | expectedType=%s sessionID=%s", expectedType, sessionID)

	// 1. Lookup por hash + tipo (evita swap entre access/refresh)
	hash := security.HashTokenHex(rawToken)
	logger.Log.Debugf("[buildTokenView] hash generado=%s", hash)

	tok, err := s.tokenRepo.FindByHashAndType(ctx, hash, expectedType)
	if err != nil {
		logger.Log.Errorf("[buildTokenView] error buscando token en repo: %v", err)
		return dto.TokenViewDTO{}, ErrTokenNotFound
	}
	if tok == nil {
		logger.Log.Debugf("[buildTokenView] token no encontrado en repo")
		return dto.TokenViewDTO{}, ErrTokenNotFound
	}
	logger.Log.Debugf("[buildTokenView] token encontrado | tokenID=%s tipo=%s status=%s sessionID(db)=%s",
		tok.TokenID, tok.TokenType, tok.Status, tok.SessionID)

	// 2. Debe pertenecer a la sesión pedida
	if tok.SessionID != sessionID {
		logger.Log.Errorf("[buildTokenView] session mismatch | esperado=%s encontrado=%s", sessionID, tok.SessionID)
		return dto.TokenViewDTO{}, ErrTokenSessionMismatch
	}

	// 3. Verificar firma/claims con go-jose (única fuente)
	claims, vErr := security.VerifyTokenRS256(rawToken)
	if vErr != nil {
		logger.Log.Debugf("[buildTokenView] error verificando firma/claims: %v", vErr)
	} else {
		logger.Log.Debugf("[buildTokenView] firma válida, claims OK")
	}

	now := time.Now().UTC()
	status := tok.Status

	// 4. Interpretación del error de verificación
	switch {
	case vErr == nil:
		logger.Log.Debugf("[buildTokenView] verificación exitosa, token válido")
	case errors.Is(vErr, security.ErrTokenExpired):
		logger.Log.Debugf("[buildTokenView] token expirado detectado")
		if status == models.TokenStatusActive {
			if err := s.tokenRepo.MarkExpired(ctx, tok.ID, now); err != nil {
				logger.Log.Errorf("[buildTokenView] error al marcar expirado en DB: %v", err)
			} else {
				logger.Log.Debugf("[buildTokenView] token marcado como expirado en DB")
			}
			status = models.TokenStatusExpired
		}
	default:
		logger.Log.Errorf("[buildTokenView] token malformado o inválido | err=%v", vErr)
		if expectedType == models.TokenTypeAccess {
			return dto.TokenViewDTO{}, ErrMalformedAccessToken
		}
		return dto.TokenViewDTO{}, ErrMalformedRefreshToken
	}

	// 5. Si DB dice revoked, eso sí es error (401)
	if status == models.TokenStatusRevoked {
		logger.Log.Debugf("[buildTokenView] token está revocado en DB | tokenID=%s", tok.TokenID)
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
	logger.Log.Debugf("[buildTokenView] detalles calculados | issuedAt=%s expiresAt=%s expiresIn=%d",
		issuedAt, expiresAt, expiresIn)

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

	logger.Log.Debugf("[buildTokenView] success | tokenID=%s tipo=%s status=%s subject=%s",
		result.TokenID, result.TokenType, result.Status, result.TokenDetail.Subject)

	return result, nil
}
