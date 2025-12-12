package handlers

import (
	"strconv"

	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/services/organic-units"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetOrganicUnitsHandler(c fiber.Ctx) error {
	// Parámetros (?page=1&page_size=20&is_deleted=true|false)
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

	rows, total, err := services.GetOrganicUnits(page, pageSize, isDeleted)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusOK).JSON(dto.OrganicUnitsListResponse{
				Data:     []dto.OrganicUnitItemDTO{},
				Total:    0,
				Page:     page,
				PageSize: pageSize,
			})
		}
		logger.Log.Error("Error obteniendo dependencias:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	out := make([]dto.OrganicUnitItemDTO, 0, len(rows))
	for _, r := range rows {
		out = append(out, mapper.ToOrganicUnitItemDTO(r.OrganicUnit, r.UsersCount))
	}

	return c.Status(fiber.StatusOK).JSON(dto.OrganicUnitsListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
