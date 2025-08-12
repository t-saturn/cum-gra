package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// Errores usados en el flujo de refresh
var (
	ErrInvalidTokenType = errors.New("invalid_token_type")
	ErrTokenExpired     = errors.New("token_expired")
	ErrSessionExpired   = errors.New("session_expired")
)

// RefreshToken genera un nuevo par access/refresh a partir de un refresh token válido.
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string, q dto.AuthRefreshQueryDTO, b dto.AuthRefreshResquestDTO) (*dto.AuthRefreshResponseDTO, error) {
	now := utils.NowUTC()

	// 0. Validación mínima
	if refreshToken == "" || q.SessionID == "" {
		return nil, ErrInvalidToken
	}

	// 1. Buscar el refresh token original en BD por HASH (no crudo)
	refreshHash := security.HashTokenHex(refreshToken)
	oldTok, err := s.tokenRepo.FindByHash(ctx, refreshHash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2. Verificar que sea realmente un refresh ACTIVO
	if oldTok.TokenType != models.TokenTypeRefresh {
		return nil, ErrInvalidTokenType
	}
	if oldTok.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}

	// 3. Verificación CRIPTO del token crudo (firma + exp)
	claims, vErr := security.VerifyTokenRS256(refreshToken)
	if vErr != nil {
		if errors.Is(vErr, security.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	// 4. Validar sesión asociada (y que coincida con la solicitada en query)
	sess, err := s.sessionRepo.FindBySessionID(ctx, oldTok.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}
	if sess.SessionID != q.SessionID {
		return nil, ErrSessionMismatch
	}
	if sess.Status != models.SessionStatusActive || !sess.IsActive {
		return nil, ErrSessionInactive
	}
	if !sess.ExpiresAt.After(now) {
		return nil, ErrSessionExpired
	}

	// Defensa adicional: subject del token debe coincidir con el user del token guardado
	if claims.Subject != "" && claims.Subject != oldTok.UserID {
		return nil, ErrInvalidToken
	}

	// 5. Revocar el access token pareado (si existe y está activo)
	if oldTok.PairedTokenID != nil {
		accessTok, err := s.tokenRepo.FindByID(ctx, oldTok.PairedTokenID.Hex())
		if err == nil && accessTok.Status == models.TokenStatusActive {
			_ = s.tokenRepo.UpdateStatus(ctx, accessTok.ID, models.TokenStatusRevoked,
				models.TokenReasonRefreshToken, oldTok.UserID, oldTok.SessionID)
		}
	}

	// 6. Revocar el refresh token antiguo
	_ = s.tokenRepo.UpdateStatus(ctx, oldTok.ID, models.TokenStatusRevoked,
		models.TokenReasonRefreshToken, oldTok.UserID, oldTok.SessionID)

	// 7. Emitir nuevo refresh token (usando DeviceInfo del body)
	refreshDuration := 7 * 24 * time.Hour
	newRefreshID, newRefreshJWT, err := s.InsertToken(ctx, oldTok.UserID, sess.SessionID, b.DeviceInfo,
		models.TokenTypeRefresh, refreshDuration, &oldTok.ID)
	if err != nil {
		return nil, err
	}
	// Relación parent → child
	_ = s.tokenRepo.AddChildToken(ctx, oldTok.ID, newRefreshID)

	// 8. Emitir nuevo access token (usando DeviceInfo del body)
	expMin, _ := strconv.Atoi(config.GetConfig().Server.JWTExpMinutes)
	accessDuration := time.Duration(expMin) * time.Minute
	newAccessID, newAccessJWT, err := s.InsertToken(ctx, oldTok.UserID, sess.SessionID, b.DeviceInfo,
		models.TokenTypeAccess, accessDuration, nil)
	if err != nil {
		return nil, err
	}

	// 9. Enlazar ambos tokens
	_ = s.tokenRepo.SetPairedTokenID(ctx, newRefreshID, newAccessID)
	_ = s.tokenRepo.SetPairedTokenID(ctx, newAccessID, newRefreshID)

	// 10. Asociar tokens a la sesión
	_ = s.sessionRepo.AddTokenToSession(ctx, sess.SessionID, newRefreshID)
	_ = s.sessionRepo.AddTokenToSession(ctx, sess.SessionID, newAccessID)

	// 11. Respuesta
	resp := &dto.AuthRefreshResponseDTO{
		AccessToken: dto.TokenDetailDTO{
			TokenID:   newAccessID.Hex(),
			Token:     newAccessJWT,
			TokenType: models.TokenTypeAccess,
			ExpiresAt: now.Add(accessDuration),
		},
		RefreshToken: dto.TokenDetailDTO{
			TokenID:   newRefreshID.Hex(),
			Token:     newRefreshJWT,
			TokenType: models.TokenTypeRefresh,
			ExpiresAt: now.Add(refreshDuration),
		},
		SessionID:    sess.SessionID,
		RefreshCount: len(oldTok.ChildTokens) + 1,
	}

	return resp, nil
}
