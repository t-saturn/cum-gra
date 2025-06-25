package handlers

import (
	"context"
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateTokenHandler(c fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id es requerido"})
	}

	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id inválido"})
	}

	token, jti, expiresAt, err := config.GenerateJWT(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo generar token"})
	}

	// Guardar en tokens_activos (simplificado)
	tokenHash := sha256.Sum256([]byte(token))
	tokenDoc := models.ActiveToken{
		TokenID:       jti,
		UserID:        uid,
		TokenHash:     hex.EncodeToString(tokenHash[:]),
		TokenType:     "access",
		IssuedAt:      time.Now(),
		ExpiresAt:     expiresAt,
		CreatedAt:     time.Now(),
		ApplicationID: "test-app", // puedes pasarlo en headers si lo necesitas
	}

	collection := config.MongoDatabase.Collection("tokens_activos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, tokenDoc); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error guardando token"})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"exp":   expiresAt,
	})
}

func ValidateTokenHandler(c fiber.Ctx) error {
	tokenStr := c.Query("token")
	appID := c.Query("app_id")

	if tokenStr == "" || appID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "token and app_id are required"})
	}

	// Verificar firma del JWT
	parsed, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET", "")), nil
	})
	if err != nil || !parsed.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token signature"})
	}

	// Extraer claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok || claims["jti"] == nil || claims["sub"] == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	jti := claims["jti"].(string)
	sub := claims["sub"].(string)

	userID, err := primitive.ObjectIDFromHex(sub)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id in token"})
	}

	// Buscar el token en la colección tokens_activos
	tokenHash := sha256.Sum256([]byte(tokenStr))
	hashHex := hex.EncodeToString(tokenHash[:])

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var found models.ActiveToken
	err = config.MongoDatabase.Collection("tokens_activos").FindOne(ctx, bson.M{
		"tokenId":       jti,
		"tokenHash":     hashHex,
		"userId":        userID,
		"applicationId": appID,
	}).Decode(&found)

	if err != nil {
		// Registro automático en tokens_invalid
		invalid := &models.InvalidToken{
			TokenID:            jti,
			TokenHash:          hashHex,
			UserID:             userID,
			ApplicationID:      appID,
			InvalidatedAt:      time.Now(),
			InvalidationReason: "invalid_token",
			InvalidatedBy:      "system",
		}
		_ = repository.CreateInvalidToken(ctx, invalid)

		config.Logger.WithField("tokenId", jti).Warn("Token inválido registrado en tokens_invalid")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not found or revoked"})
	}

	// Token válido
	return c.JSON(fiber.Map{
		"message": "Token is valid",
		"tokenId": jti,
		"userId":  userID.Hex(),
		"exp":     found.ExpiresAt,
	})
}
