package handlers

import (
	"strconv"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/mapper"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetOrganicUnitsHandler(c fiber.Ctx) error {
	db := config.DB

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

	var total int64
	base := db.Model(&models.OrganicUnit{}).Where("is_deleted = ?", isDeleted)
	if err := base.Count(&total).Error; err != nil {
		logger.Log.Error("Error contando dependencias:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	type row struct {
		models.OrganicUnit
		UsersCount int64 `gorm:"column:users_count"`
	}
	var rows []row

	q := db.
		Table("organic_units ou").
		Select(`
			ou.*,
			COUNT(u.id) AS users_count
		`).
		Joins(`
			LEFT JOIN users u
				ON u.organic_unit_id = ou.id
				AND u.is_deleted = FALSE
		`).
		Where("ou.is_deleted = ?", isDeleted).
		Group("ou.id").
		Order("ou.created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := q.Scan(&rows).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusOK).JSON(dto.OrganicUnitsListResponse{
				Data:     []dto.OrganicUnitItemDTO{},
				Total:    0,
				Page:     page,
				PageSize: pageSize,
			})
		}
		logger.Log.Error("Error obteniendo dependencias:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
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
