package handlers

import (
	"strconv"

	"server/internal/dto"
	services "server/internal/services/organic-units"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetOrganicUnitsHandler(c fiber.Ctx) error {
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

	isDeleted := false
	if v := c.Query("is_deleted"); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			isDeleted = b
		}
	}

	var parentID *string
	if v := c.Query("parent_id"); v != "" {
		parentID = &v
	}

	result, err := services.GetOrganicUnits(page, pageSize, isDeleted, parentID)
	if err != nil {
		logger.Log.Error("Error obteniendo unidades orgánicas:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}