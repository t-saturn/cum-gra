package handlers

import (
	"context"
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
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
