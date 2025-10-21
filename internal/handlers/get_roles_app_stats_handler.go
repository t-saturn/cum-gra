package handlers

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"

	"github.com/gofiber/fiber/v3"
)

func GetRolesAppStatsHandler(c fiber.Ctx) error {
	db := config.DB
	var totalRoles, activeRoles, adminRoles, assignedUsers int64

	if err := db.Model(&models.ApplicationRole{}).
		Count(&totalRoles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to count total roles",
		})
	}

	if err := db.Model(&models.ApplicationRole{}).
		Where("is_deleted = FALSE").
		Count(&activeRoles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to count active roles",
		})
	}

	if err := db.Model(&models.ApplicationRole{}).
		Where("is_deleted = FALSE AND name ILIKE '%admin%'").
		Count(&adminRoles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to count admin roles",
		})
	}

	if err := db.Table("user_application_roles uar").
		Joins("JOIN application_roles ar ON ar.id = uar.application_role_id").
		Where("uar.is_deleted = FALSE AND ar.is_deleted = FALSE").
		Distinct("uar.user_id").
		Count(&assignedUsers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to count assigned users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.RolesAppStatsResponse{
		TotalRoles:    totalRoles,
		ActiveRoles:   activeRoles,
		AdminRoles:    adminRoles,
		AssignedUsers: assignedUsers,
	})
}
