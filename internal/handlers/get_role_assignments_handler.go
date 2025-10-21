package handlers

import (
	"strconv"

	"central-user-manager/internal/dto"
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetRoleAssignmentsHandler(c fiber.Ctx) error {
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
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid is_deleted (use true|false)",
			})
		}
	}

	assignments, err := services.GetRoleAssignments(page, pageSize, isDeleted)
	if err != nil {
		logger.Log.Error("Error obteniendo asignaciones de roles:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Error interno del servidor",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.RolesAssignmentsResponseDTO{
		Assignments: assignments,
	})
}
