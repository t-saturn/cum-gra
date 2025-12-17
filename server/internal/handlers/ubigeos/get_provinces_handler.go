package handlers

import (
	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetProvincesByDepartmentHandler(c fiber.Ctx) error {
	department := c.Query("department")

	result, err := services.GetProvincesByDepartment(department)
	if err != nil {
		if err.Error() == "departamento requerido" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error obteniendo provincias:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}