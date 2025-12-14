package handlers

import (
	"strconv"

	"server/internal/dto"
	services "server/internal/services/users"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUsersHandler(c fiber.Ctx) error {
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

	var status *string
	if v := c.Query("status"); v != "" {
		status = &v
	}

	var organicUnitID *string
	if v := c.Query("organic_unit_id"); v != "" {
		organicUnitID = &v
	}

	var positionID *string
	if v := c.Query("position_id"); v != "" {
		positionID = &v
	}

	result, err := services.GetUsers(page, pageSize, isDeleted, status, organicUnitID, positionID)
	if err != nil {
		logger.Log.Error("Error obteniendo usuarios:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}