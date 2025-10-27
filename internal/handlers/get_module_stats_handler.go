package handlers

import (
	"central-user-manager/internal/dto"
	"central-user-manager/internal/services"

	"github.com/gofiber/fiber/v3"
)

func GetModulesStatsHandler(c fiber.Ctx) error {
	stats, err := services.GetModulesStats()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}
