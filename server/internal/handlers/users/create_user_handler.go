package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/users"
	"server/pkg/logger"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func CreateUserHandler(c fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Datos inválidos"})
	}

	// Obtener email del contexto
	email, ok := c.Locals("email").(string)
	if !ok || email == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Usuario no autenticado"})
	}

	// Buscar el usuario en la BD por email
	var user models.User
	db := config.DB
	if err := db.Where("email = ? AND is_deleted = FALSE", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Usuario no encontrado en el sistema"})
	}

	// Obtener access token del header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Token no proporcionado"})
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	if accessToken == authHeader {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Formato de token inválido"})
	}

	result, err := services.CreateUser(req, user.ID, accessToken)
	if err != nil {
		// Manejo de errores específicos de Keycloak
		if strings.Contains(err.Error(), "keycloak") {
			return c.Status(fiber.StatusBadGateway).
				JSON(dto.ErrorResponse{Error: "Error de comunicación con el sistema de autenticación"})
		}

		if err.Error() == "ya existe un usuario con este email" ||
			err.Error() == "ya existe un usuario con este DNI" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}

		if strings.Contains(err.Error(), "no encontrada") ||
			strings.Contains(err.Error(), "inválido") {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}

		logger.Log.Error("Error creando usuario:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}