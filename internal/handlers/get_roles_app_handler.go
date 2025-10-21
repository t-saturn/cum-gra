package handlers

import (
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetRolesAppHandler(c fiber.Ctx) error {
	page := c.Query("page", "")
	pageSize := c.Query("page_size", "")
	isDeleted := c.Query("is_deleted", "")

	resp, err := services.GetRolesApp(page, pageSize, isDeleted)
	if err != nil {
		logger.Log.Error("GetRolesAppHandler.service", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
