package handlers

import (
	"server/internal/dto"
	"server/internal/services/users"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUsersStatsHandler(c fiber.Ctx) error {
	total, active, suspended, newLastMonth, err := services.GetUserStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de usuarios:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.UsersStatsResponse{
		TotalUsers:        total,
		ActiveUsers:       active,
		SuspendedUsers:    suspended,
		NewUsersLastMonth: newLastMonth,
	})
}
