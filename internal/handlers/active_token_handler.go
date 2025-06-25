package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/repository"
)

func GetAllActiveTokensHandler(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tokens, err := repository.GetAllActiveTokens(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch active tokens"})
	}
	return c.JSON(tokens)
}
