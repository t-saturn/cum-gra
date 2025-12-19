package handlers

import (
	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetDistrictsByProvinceHandler(c fiber.Ctx) error {
	department := c.Query("department")
	province := c.Query("province")

	result, err := services.GetDistrictsByProvince(department, province)
	if err != nil {
		if err.Error() == "departamento requerido" || err.Error() == "provincia requerida" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error obteniendo distritos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}