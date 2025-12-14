package handlers

import (
	"strconv"

	"server/internal/dto"
	services "server/internal/services/positions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetStructuralPositionsHandler(c fiber.Ctx) error {
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

	var level *int
	if v := c.Query("level"); v != "" {
		if l, err := strconv.Atoi(v); err == nil && l > 0 {
			level = &l
		}
	}

	result, err := services.GetStructuralPositions(page, pageSize, isDeleted, level)
	if err != nil {
		logger.Log.Error("Error obteniendo posiciones estructurales:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}