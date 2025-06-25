package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repository"
)

func CreateInvalidTokenHandler(c fiber.Ctx) error {
	var token models.InvalidToken
	if err := c.Bind().Body(&token); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token.InvalidatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repository.CreateInvalidToken(ctx, &token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error saving invalid token"})
	}

	return c.JSON(fiber.Map{"message": "Invalid token saved"})
}

func GetAllInvalidTokensHandler(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tokens, err := repository.GetAllInvalidTokens(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error retrieving tokens"})
	}

	return c.JSON(tokens)
}
