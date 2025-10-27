package handlers

import (
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetRolesAppStatsHandler(c fiber.Ctx) error {
	resp, err := services.GetRolesAppStats()
	if err != nil {
		logger.Log.Error("GetRolesAppStatsHandler.service", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch roles app stats",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
