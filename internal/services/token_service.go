package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateAndStoreToken(userID string, appID, appURL string, device models.DeviceInfo) (string, error) {
	tokenStr, jti, exp, err := config.GenerateJWT(userID, "access")
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

func RefreshAccessToken(refreshToken string) (string, error) {
	// 1. Validar JWT
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(config.GetEnv("JWT_SECRET", "")), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("refresh token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["jti"] == nil || claims["sub"] == nil {
		return "", fmt.Errorf("token incompleto")
	}

	jti := claims["jti"].(string)
	userID := claims["sub"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 2. Buscar en Mongo la info del token
	tokenDoc, err := repository.FindTokenByJTI(ctx, jti)
	if err != nil || tokenDoc.Status != "active" || tokenDoc.TokenType != "refresh" {
		return "", fmt.Errorf("refresh token inválido o inactivo")
	}
	if tokenDoc.RefreshCount >= tokenDoc.MaxRefreshCount {
		return "", fmt.Errorf("refresh token agotado")
	}

	// 3. Generar nuevo access_token
	newAccessToken, newJti, exp, err := config.GenerateJWT(userID, "access")
	if err != nil {
		return "", fmt.Errorf("error generando nuevo token")
	}

	// 4. Guardar nuevo token como access
	hash := sha256.Sum256([]byte(newAccessToken))
	accessToken := &models.Token{
		TokenID:        newJti,
		TokenHash:      hex.EncodeToString(hash[:]),
		UserID:         tokenDoc.UserID,
		SessionID:      tokenDoc.SessionID,
		Status:         "active",
		TokenType:      "access",
		IssuedAt:       time.Now(),
		ExpiresAt:      exp,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ApplicationID:  tokenDoc.ApplicationID,
		ApplicationURL: tokenDoc.ApplicationURL,
		DeviceInfo:     tokenDoc.DeviceInfo,
	}

	if err := repository.InsertToken(ctx, accessToken); err != nil {
		return "", fmt.Errorf("error guardando nuevo token")
	}

	// 5. Incrementar refreshCount
	_ = repository.IncrementRefreshCount(ctx, tokenDoc.TokenID)

	return newAccessToken, nil
}
