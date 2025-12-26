package handlers

import (
	"server/internal/dto"
	services "server/internal/services/users"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetAllUsersHandler(c fiber.Ctx) error {
	onlyActive := true
	if v := c.Query("only_active"); v == "false" {
		onlyActive = false
	}

	result, err := services.GetAllUsers(onlyActive)
	if err != nil {
		logger.Log.Error("Error obteniendo todos los usuarios:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}