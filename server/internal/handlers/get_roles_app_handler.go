package handlers

import (
	"strconv"
	"strings"

	"central-user-manager/internal/dto"
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetRolesAppHandler(c fiber.Ctx) error {
	pageStr := c.Query("page", "")
	pageSizeStr := c.Query("page_size", "")
	isDeletedStr := strings.ToLower(strings.TrimSpace(c.Query("is_deleted", "")))

	page := 1
	if n, err := strconv.Atoi(strings.TrimSpace(pageStr)); err == nil && n > 0 {
		page = n
	}
	pageSize := 20
	if n, err := strconv.Atoi(strings.TrimSpace(pageSizeStr)); err == nil && n > 0 && n <= 200 {
		pageSize = n
	}

	var isDeletedPtr *bool
	switch isDeletedStr {
	case "true", "1", "t":
		v := true
		isDeletedPtr = &v
	case "false", "0", "f":
		v := false
		isDeletedPtr = &v
	}

	resp, err := services.GetRolesAppMatrix(services.GetRolesAppOptions{
		Page:      page,
		PageSize:  pageSize,
		IsDeleted: isDeletedPtr,
	})
	if err != nil {
		logger.Log.Errorf("GetRolesAppHandler.service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error interno del servidor",
		})
	}

	if resp == nil {
		return c.JSON(dto.RolesAppResponseDTO{
			Data:     []dto.RoleAppModulesItemDTO{},
			Total:    0,
			Page:     page,
			PageSize: pageSize,
		})
	}
	return c.JSON(resp)
}
