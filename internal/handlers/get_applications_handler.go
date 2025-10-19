package handlers

import (
	"strconv"
	"strings"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/mapper"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetApplicationsHandler(c fiber.Ctx) error {
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

	adminRoleName := c.Query("admin_role_name", "admin")
	adminRolePrefix := strings.TrimSpace(adminRoleName)

	var total int64
	base := db.Model(&models.Application{}).Where("is_deleted = ?", isDeleted)
	if err := base.Count(&total).Error; err != nil {
		logger.Log.Error("Error contando aplicaciones:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var apps []models.Application
	if err := db.
		Where("is_deleted = ?", isDeleted).
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&apps).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusOK).JSON(dto.ApplicationsListResponse{
				Data:     []dto.ApplicationDTO{},
				Total:    0,
				Page:     page,
				PageSize: pageSize,
			})
		}
		logger.Log.Error("Error obteniendo aplicaciones:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	if len(apps) == 0 {
		return c.Status(fiber.StatusOK).JSON(dto.ApplicationsListResponse{
			Data:     []dto.ApplicationDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
	}

	appIDs := make([]uuid.UUID, 0, len(apps))
	for _, a := range apps {
		appIDs = append(appIDs, a.ID)
	}

	type countRow struct {
		ApplicationID uuid.UUID `gorm:"column:application_id"`
		UsersCount    int64     `gorm:"column:users_count"`
	}
	var counts []countRow

	if err := db.
		Table("user_application_roles").
		Select("application_id, COUNT(DISTINCT user_id) AS users_count").
		Where("is_deleted = FALSE AND revoked_at IS NULL AND application_id IN ?", appIDs).
		Group("application_id").
		Scan(&counts).Error; err != nil {

		logger.Log.Error("Error contando usuarios por aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	usersCountByApp := make(map[uuid.UUID]int64, len(counts))
	for _, r := range counts {
		usersCountByApp[r.ApplicationID] = r.UsersCount
	}

	type adminRow struct {
		ApplicationID uuid.UUID `gorm:"column:application_id"`
		UserID        uuid.UUID `gorm:"column:user_id"`
		Email         string    `gorm:"column:email"`
		FirstName     *string   `gorm:"column:first_name"`
		LastName      *string   `gorm:"column:last_name"`
		DNI           string    `gorm:"column:dni"`
	}

	var admins []adminRow

	pattern := adminRolePrefix + "%"

	if err := db.
		Table("application_roles ar").
		Select(`
			ar.application_id,
			uar.user_id,
			u.email,
			u.first_name,
			u.last_name,
			u.dni
		`).
		Joins(`JOIN user_application_roles uar
					ON uar.application_role_id = ar.id
					AND uar.is_deleted = FALSE
					AND uar.revoked_at IS NULL`).
		Joins(`JOIN users u
					ON u.id = uar.user_id
					AND u.is_deleted = FALSE`).
		Where("ar.is_deleted = FALSE AND ar.application_id IN ? AND ar.name ILIKE ?", appIDs, pattern).
		Scan(&admins).Error; err != nil {

		logger.Log.Error("Error obteniendo administradores por aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	adminsByApp := make(map[uuid.UUID][]dto.AdminUserDTO)
	for _, a := range admins {
		u := models.User{
			ID:        a.UserID,
			Email:     a.Email,
			FirstName: a.FirstName,
			LastName:  a.LastName,
			DNI:       a.DNI,
		}
		adminsByApp[a.ApplicationID] = append(adminsByApp[a.ApplicationID], mapper.ToAdminUserDTO(u))
	}

	out := make([]dto.ApplicationDTO, 0, len(apps))
	for _, a := range apps {
		adminList := adminsByApp[a.ID]
		usersCount := usersCountByApp[a.ID]
		out = append(out, mapper.ToApplicationDTO(a, adminList, usersCount))
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApplicationsListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
