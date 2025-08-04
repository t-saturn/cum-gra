package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// Error variables for RefreshToken flow
var (
	ErrInvalidTokenType = errors.New("invalid_token_type")
	ErrTokenExpired     = errors.New("token_expired")
	ErrSessionExpired   = errors.New("session_expired")
)

// RefreshToken genera un nuevo par access/refresh a partir de un refresh token válido.
func (s *AuthService) RefreshToken(ctx context.Context, input dto.AuthRefreshRequestDTO) (*dto.AuthRefreshResponseDTO, error) {
	now := utils.NowUTC()

	// 1) Buscar el refresh token original en BD
	oldTok, err := s.tokenRepo.FindByHash(ctx, input.RefreshToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2) Verificar que sea realmente un refresh activo y no expirado
	if oldTok.TokenType != models.TokenTypeRefresh || oldTok.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}
	if !oldTok.ExpiresAt.After(now) {
		return nil, ErrTokenExpired
	}

	// 3) Buscar y validar sesión asociada
	sess, err := s.sessionRepo.FindBySessionID(ctx, oldTok.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}
	if sess.Status != models.SessionStatusActive || !sess.IsActive {
		return nil, ErrSessionInactive
	}
	if !sess.ExpiresAt.After(now) {
		return nil, ErrSessionExpired
	}

	// 4) Revocar el refresh antiguo
	reason := models.TokenReasonRefreshToken
	revokedBy := oldTok.UserID
	revokedByApp := oldTok.SessionID

	if err := s.tokenRepo.UpdateStatus(ctx, oldTok.ID, models.TokenStatusRevoked, reason, revokedBy, revokedByApp); err != nil {
		return nil, err
	}

	// 5) Crear y persistir el nuevo refresh token
	refreshID, refreshJWT, err := s.InsertToken(ctx, oldTok.UserID, sess.SessionID, input.DeviceInfo, models.TokenTypeRefresh, 7*24*time.Hour, &oldTok.ID)
	if err != nil {
		return nil, err
	}
	// registrar relación parent → child
	_ = s.tokenRepo.AddChildToken(ctx, oldTok.ID, refreshID)

	// 6) Crear y persistir el nuevo access token
	expMin, _ := strconv.Atoi(config.GetConfig().Server.JWTExpMinutes)
	accessID, accessJWT, err := s.InsertToken(ctx, oldTok.UserID, sess.SessionID, input.DeviceInfo, models.TokenTypeAccess, time.Duration(expMin)*time.Minute, nil)
	if err != nil {
		return nil, err
	}
	// enlazar ambos tokens
	_ = s.tokenRepo.SetPairedTokenID(ctx, refreshID, accessID)
	_ = s.tokenRepo.SetPairedTokenID(ctx, accessID, refreshID)

	// 7) Asociar tokens a la sesión
	_ = s.sessionRepo.AddTokenToSession(ctx, sess.SessionID, refreshID)
	_ = s.sessionRepo.AddTokenToSession(ctx, sess.SessionID, accessID)

	// 8) Construir respuesta
	resp := &dto.AuthRefreshResponseDTO{
		AccessToken: dto.TokenDetailDTO{
			TokenID:   accessID.Hex(),
			Token:     accessJWT,
			TokenType: models.TokenTypeAccess,
			ExpiresAt: now.Add(time.Duration(expMin) * time.Minute),
		},
		RefreshToken: dto.TokenDetailDTO{
			TokenID:   refreshID.Hex(),
			Token:     refreshJWT,
			TokenType: models.TokenTypeRefresh,
			ExpiresAt: now.Add(7 * 24 * time.Hour),
		},
		SessionID:    sess.SessionID,
		RefreshCount: len(oldTok.ChildTokens) + 1,
	}

	return resp, nil
}
