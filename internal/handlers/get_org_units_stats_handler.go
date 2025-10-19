package handlers

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetOrganicUnitsStatsHandler(c fiber.Ctx) error {
	db := config.DB

	var totalUnits int64
	if err := db.Model(&models.OrganicUnit{}).
		Count(&totalUnits).Error; err != nil {
		logger.Log.Error("Error contando total de unidades orgánicas:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var activeUnits int64
	if err := db.Model(&models.OrganicUnit{}).
		Where("is_deleted = ? AND is_active = ?", false, true).
		Count(&activeUnits).Error; err != nil {
		logger.Log.Error("Error contando unidades orgánicas activas:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var deletedUnits int64
	if err := db.Model(&models.OrganicUnit{}).
		Where("is_deleted = ?", true).
		Count(&deletedUnits).Error; err != nil {
		logger.Log.Error("Error contando unidades orgánicas eliminadas:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var totalEmployees int64
	if err := db.Table("users u").
		Joins("JOIN organic_units ou ON ou.id = u.organic_unit_id").
		Where("u.is_deleted = FALSE AND ou.is_deleted = FALSE").
		Count(&totalEmployees).Error; err != nil {
		logger.Log.Error("Error contando empleados:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.OrganicUnitsStatsResponse{
		TotalOrganicUnits:   totalUnits,
		ActiveOrganicUnits:  activeUnits,
		DeletedOrganicUnits: deletedUnits,
		TotalEmployees:      totalEmployees,
	})
}
