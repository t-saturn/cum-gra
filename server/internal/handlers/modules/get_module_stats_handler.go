package handlers

import (
	"server/internal/dto"
	"server/internal/services/modules"

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
