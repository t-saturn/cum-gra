package services

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenService interface {
	GenerateAndStoreToken(ctx context.Context, userID, applicationID string, deviceInfo models.DeviceInfo) (*models.Token, string, error)
}

type tokenService struct {
	repo repositories.TokenRepository
}

func NewTokenService(repo repositories.TokenRepository) TokenService {
	return &tokenService{repo: repo}
}

func (s *tokenService) GenerateAndStoreToken(ctx context.Context, userID, applicationID string, deviceInfo models.DeviceInfo) (*models.Token, string, error) {
	tokenStr, jti, exp, err := config.GenerateJWT(userID, "access")
	if err != nil {
		return nil, "", err
	}

	uid, _ := primitive.ObjectIDFromHex(userID)
	token := &models.Token{
		TokenID:         jti,
		TokenHash:       tokenStr, // puedes hashearlo si quieres
		UserID:          uid,
		Status:          "active",
		TokenType:       "access",
		IssuedAt:        time.Now(),
		ExpiresAt:       exp,
		ApplicationID:   applicationID,
		DeviceInfo:      deviceInfo,
		MaxRefreshCount: 5,
		RefreshCount:    0,
	}

	err = s.repo.Insert(ctx, token)
	if err != nil {
		return nil, "", err
	}

	return token, tokenStr, nil
}
