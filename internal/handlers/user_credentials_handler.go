package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repository"
)

func InsertCredentialHandler(c fiber.Ctx) error {
	var input models.UserCredential
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato inválido"})
	}

	// Generar instancia limpia
	cred := models.NewUserCredential(
		input.Email,
		input.PasswordHash,
		input.ApplicationID,
		input.ApplicationURL,
		input.DeviceInfo,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repository.CreateUserCredential(ctx, cred); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al insertar"})
	}

	return c.JSON(fiber.Map{"message": "Credencial registrada"})
}

func GetAllCredentialsHandler(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds, err := repository.GetAllUserCredentials(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al obtener credenciales"})
	}

	return c.JSON(creds)
}
