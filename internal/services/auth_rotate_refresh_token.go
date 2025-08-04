package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/security"
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

	// 1. Buscar token por su valor
	oldTok, err := s.tokenRepo.FindByHash(ctx, input.RefreshToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// Validaciones básicas del token
	if oldTok.TokenType != models.TokenTypeRefresh {
		return nil, ErrInvalidTokenType
	}
	if oldTok.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}
	if !oldTok.ExpiresAt.After(now) {
		return nil, ErrTokenExpired
	}

	// 2. Verificar la sesión asociada
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

	// 3. Revocar el refresh token anterior
	if err := s.tokenRepo.UpdateStatus(ctx, oldTok.ID, models.TokenStatusRevoked, &now, nil); err != nil {
		return nil, err
	}

	// 4. Generar nuevo access token (1 minuto)
	accessJWT, err := security.GenerateToken(oldTok.UserID, time.Minute)
	if err != nil {
		return nil, err
	}
	accessModel := &models.Token{
		TokenID:   uuid.New().String(),
		TokenHash: accessJWT,
		UserID:    oldTok.UserID,
		SessionID: oldTok.SessionID,
		Status:    models.TokenStatusActive,
		TokenType: models.TokenTypeAccess,
		IssuedAt:  now,
		ExpiresAt: now.Add(time.Minute),
		CreatedAt: now,
		UpdatedAt: now,
		DeviceInfo: models.DeviceInfo{
			UserAgent:      input.DeviceInfo.UserAgent,
			IP:             input.DeviceInfo.IP,
			DeviceID:       input.DeviceInfo.DeviceID,
			BrowserName:    input.DeviceInfo.BrowserName,
			BrowserVersion: input.DeviceInfo.BrowserVersion,
			OS:             input.DeviceInfo.OS,
			OSVersion:      input.DeviceInfo.OSVersion,
			DeviceType:     input.DeviceInfo.DeviceType,
			Timezone:       input.DeviceInfo.Timezone,
			Language:       input.DeviceInfo.Language,
		},
	}
	if input.DeviceInfo.Location != nil {
		accessModel.DeviceInfo.Location = &models.LocationDetail{
			Country:     input.DeviceInfo.Location.Country,
			CountryCode: input.DeviceInfo.Location.CountryCode,
			Region:      input.DeviceInfo.Location.Region,
			City:        input.DeviceInfo.Location.City,
			Coordinates: models.Coordinates{
				input.DeviceInfo.Location.Coordinates[0],
				input.DeviceInfo.Location.Coordinates[1],
			},
			ISP:          input.DeviceInfo.Location.ISP,
			Organization: input.DeviceInfo.Location.Organization,
		}
	}

	accessOID, err := s.tokenRepo.Insert(ctx, accessModel)
	if err != nil {
		return nil, err
	}

	// 5. Generar nuevo refresh token (15 minutos)
	refreshJWT, err := security.GenerateToken(oldTok.UserID, 15*time.Minute)
	if err != nil {
		return nil, err
	}
	refreshModel := &models.Token{
		TokenID:       uuid.New().String(),
		TokenHash:     refreshJWT,
		UserID:        oldTok.UserID,
		SessionID:     oldTok.SessionID,
		Status:        models.TokenStatusActive,
		TokenType:     models.TokenTypeRefresh,
		IssuedAt:      now,
		ExpiresAt:     now.Add(15 * time.Minute),
		CreatedAt:     now,
		UpdatedAt:     now,
		ParentTokenID: &accessOID,
		DeviceInfo: models.DeviceInfo{
			UserAgent:      input.DeviceInfo.UserAgent,
			IP:             input.DeviceInfo.IP,
			DeviceID:       input.DeviceInfo.DeviceID,
			BrowserName:    input.DeviceInfo.BrowserName,
			BrowserVersion: input.DeviceInfo.BrowserVersion,
			OS:             input.DeviceInfo.OS,
			OSVersion:      input.DeviceInfo.OSVersion,
			DeviceType:     input.DeviceInfo.DeviceType,
			Timezone:       input.DeviceInfo.Timezone,
			Language:       input.DeviceInfo.Language,
		},
	}
	if input.DeviceInfo.Location != nil {
		refreshModel.DeviceInfo.Location = &models.LocationDetail{
			Country:     input.DeviceInfo.Location.Country,
			CountryCode: input.DeviceInfo.Location.CountryCode,
			Region:      input.DeviceInfo.Location.Region,
			City:        input.DeviceInfo.Location.City,
			Coordinates: models.Coordinates{
				input.DeviceInfo.Location.Coordinates[0],
				input.DeviceInfo.Location.Coordinates[1],
			},
			ISP:          input.DeviceInfo.Location.ISP,
			Organization: input.DeviceInfo.Location.Organization,
		}
	}

	refreshOID, err := s.tokenRepo.Insert(ctx, refreshModel)
	if err != nil {
		return nil, err
	}

	// 6. Construir DTO de respuesta
	resp := &dto.AuthRefreshResponseDTO{
		AccessToken: dto.TokenDetailDTO{
			TokenID:   accessOID.Hex(),
			Token:     accessJWT,
			TokenType: models.TokenTypeAccess,
			ExpiresAt: accessModel.ExpiresAt,
		},
		RefreshToken: dto.TokenDetailDTO{
			TokenID:   refreshOID.Hex(),
			Token:     refreshJWT,
			TokenType: models.TokenTypeRefresh,
			ExpiresAt: refreshModel.ExpiresAt,
		},
		SessionID: oldTok.SessionID,
	}

	return resp, nil
}
