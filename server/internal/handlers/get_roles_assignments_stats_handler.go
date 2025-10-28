package handlers

import (
	"strconv"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"

	"github.com/gofiber/fiber/v3"
)

func GetRoleAssignmentsStatsHandler(c fiber.Ctx) error {
	db := config.DB

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

	var totalUsers int64
	var adminUsers int64
	var usersWithRoles int64

	if err := db.Model(&models.User{}).
		Where("is_deleted = ?", isDeleted).
		Count(&totalUsers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to count total users",
		})
	}

	if err := db.Table("users u").
		Joins("JOIN user_application_roles uar ON uar.user_id = u.id AND uar.is_deleted = FALSE").
		Joins("JOIN application_roles ar ON ar.id = uar.application_role_id AND ar.is_deleted = FALSE").
		Where("u.is_deleted = ? AND ar.name ILIKE '%admin%'", isDeleted).
		Distinct("u.id").
		Count(&adminUsers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to count admin users",
		})
	}

	if err := db.Table("users u").
		Joins("JOIN user_application_roles uar ON uar.user_id = u.id AND uar.is_deleted = FALSE").
		Joins("JOIN application_roles ar ON ar.id = uar.application_role_id AND ar.is_deleted = FALSE").
		Where("u.is_deleted = ?", isDeleted).
		Distinct("u.id").
		Count(&usersWithRoles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to count users with roles",
		})
	}

	usersWithoutRoles := max(totalUsers-usersWithRoles, 0)

	return c.Status(fiber.StatusOK).JSON(dto.RolesAssigmentsResponseResponse{
		TotalUsers:        totalUsers,
		AdminUsers:        adminUsers,
		UsersWithRoles:    usersWithRoles,
		UsersWithoutRoles: usersWithoutRoles,
	})
}
