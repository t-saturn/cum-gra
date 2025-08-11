package services

import (
	"context"
	"errors"
	"fmt"
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

	// 4) Buscar y revocar el access token asociado al refresh token actual
	if oldTok.PairedTokenID != nil {
		fmt.Printf("DEBUG: Intentando buscar access_token con ID: %s\n", oldTok.PairedTokenID.Hex())
		accessTok, err := s.tokenRepo.FindByID(ctx, oldTok.PairedTokenID.Hex())
		if err != nil {
			fmt.Printf("DEBUG: Error al buscar access_token: %v\n", err)
		} else if accessTok.Status == models.TokenStatusActive {
			fmt.Printf("DEBUG: Revocando access_token con ID: %s\n", accessTok.ID.Hex())
			reason := models.TokenReasonRefreshToken
			revokedBy := oldTok.UserID
			revokedByApp := oldTok.SessionID
			if err := s.tokenRepo.UpdateStatus(ctx, accessTok.ID, models.TokenStatusRevoked, reason, revokedBy, revokedByApp); err != nil {
				fmt.Printf("DEBUG: Error al revocar access_token: %v\n", err)
				return nil, err
			}
			fmt.Printf("DEBUG: access_token %s revocado exitosamente\n", accessTok.ID.Hex())
		} else {
			fmt.Printf("DEBUG: access_token %s ya no está activo (estado: %s)\n", accessTok.ID.Hex(), accessTok.Status)
		}
	} else {
		fmt.Println("DEBUG: No hay PairedTokenID asociado al refresh_token")
	}

	// 5) Revocar el refresh token antiguo
	reason := models.TokenReasonRefreshToken
	revokedBy := oldTok.UserID
	revokedByApp := oldTok.SessionID

	fmt.Printf("DEBUG: Revocando refresh_token con ID: %s\n", oldTok.ID.Hex())
	if err := s.tokenRepo.UpdateStatus(ctx, oldTok.ID, models.TokenStatusRevoked, reason, revokedBy, revokedByApp); err != nil {
		fmt.Printf("DEBUG: Error al revocar refresh_token: %v\n", err)
		return nil, err
	}
	fmt.Printf("DEBUG: refresh_token %s revocado exitosamente\n", oldTok.ID.Hex())

	// 6) Crear y persistir el nuevo refresh token
	refreshID, refreshJWT, err := s.InsertToken(ctx, oldTok.UserID, sess.SessionID, input.DeviceInfo, models.TokenTypeRefresh, 7*24*time.Hour, &oldTok.ID)
	if err != nil {
		return nil, err
	}
	// registrar relación parent → child
	_ = s.tokenRepo.AddChildToken(ctx, oldTok.ID, refreshID)

	// 7) Crear y persistir el nuevo access token
	expMin, _ := strconv.Atoi(config.GetConfig().Server.JWTExpMinutes)
	accessID, accessJWT, err := s.InsertToken(ctx, oldTok.UserID, sess.SessionID, input.DeviceInfo, models.TokenTypeAccess, time.Duration(expMin)*time.Minute, nil)
	if err != nil {
		return nil, err
	}
	// enlazar ambos tokens
	_ = s.tokenRepo.SetPairedTokenID(ctx, refreshID, accessID)
	_ = s.tokenRepo.SetPairedTokenID(ctx, accessID, refreshID)

	// 8) Asociar tokens a la sesión
	_ = s.sessionRepo.AddTokenToSession(ctx, sess.SessionID, refreshID)
	_ = s.sessionRepo.AddTokenToSession(ctx, sess.SessionID, accessID)

	// 9) Construir respuesta
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
