package handlers

import (
	"strconv"

	"server/internal/dto"
	services "server/internal/services/user-application-roles"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUserApplicationRolesHandler(c fiber.Ctx) error {
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

	var userID *string
	if v := c.Query("user_id"); v != "" {
		userID = &v
	}

	var applicationID *string
	if v := c.Query("application_id"); v != "" {
		applicationID = &v
	}

	var isRevoked *string
	if v := c.Query("is_revoked"); v != "" {
		isRevoked = &v
	}

	result, err := services.GetUserApplicationRoles(page, pageSize, isDeleted, userID, applicationID, isRevoked)
	if err != nil {
		logger.Log.Error("Error obteniendo asignaciones de roles:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}