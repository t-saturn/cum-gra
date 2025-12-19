package handlers

import (
	"strconv"

	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUbigeosHandler(c fiber.Ctx) error {
	page := 1
	pageSize := 20

	if v := c.Query("page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			page = n
		}
	}

	if v := c.Query("page_size"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 200 {
			pageSize = n
		}
	}

	var department *string
	if v := c.Query("department"); v != "" {
		department = &v
	}

	var province *string
	if v := c.Query("province"); v != "" {
		province = &v
	}

	var district *string
	if v := c.Query("district"); v != "" {
		district = &v
	}

	result, err := services.GetUbigeos(page, pageSize, department, province, district)
	if err != nil {
		logger.Log.Error("Error obteniendo ubigeos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}