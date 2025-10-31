package handlers

import (
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetRoleAssignmentsHandler(c fiber.Ctx) error {
	page := c.Query("page", "")
	pageSize := c.Query("page_size", "")
	isDeleted := c.Query("is_deleted", "")

	resp, err := services.GetRoleAssignments(page, pageSize, isDeleted)
	if err != nil {
		logger.Log.Error("GetRoleAssignmentsHandler.service", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
