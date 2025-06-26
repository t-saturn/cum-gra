package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateAndStoreToken(userID string, appID, appURL string, device models.DeviceInfo) (string, error) {
	tokenStr, jti, exp, err := config.GenerateJWT(userID)
	if err != nil {
		return "", err
	}

	userObjID, _ := primitive.ObjectIDFromHex(userID)
	hash := sha256.Sum256([]byte(tokenStr))

	tokenDoc := &models.Token{
		TokenID:        jti,
		TokenHash:      hex.EncodeToString(hash[:]),
		UserID:         userObjID,
		Status:         "active",
		TokenType:      "access",
		IssuedAt:       time.Now(),
		ExpiresAt:      exp,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ApplicationID:  appID,
		ApplicationURL: appURL,
		DeviceInfo:     device,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repository.InsertToken(ctx, tokenDoc); err != nil {
		return "", err
	}

	return tokenStr, nil
}
