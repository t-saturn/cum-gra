package handlers

import (
	"central-user-manager/internal/dto"
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetPositionsStatsHandler(c fiber.Ctx) error {
	total, active, deleted, assigned, err := services.GetPositionsStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de cargos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.StructuralPositionsStatsResponse{
		TotalPositions:    total,
		ActivePositions:   active,
		DeletedPositions:  deleted,
		AssignedEmployees: assigned,
	})
}
