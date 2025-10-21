package services

import (
	"strings"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/mapper"
	"central-user-manager/internal/models"

	"github.com/google/uuid"
)

type CountRow struct {
	ApplicationID uuid.UUID `gorm:"column:application_id"`
	UsersCount    int64     `gorm:"column:users_count"`
}

type AdminRow struct {
	ApplicationID uuid.UUID `gorm:"column:application_id"`
	UserID        uuid.UUID `gorm:"column:user_id"`
	Email         string    `gorm:"column:email"`
	FirstName     *string   `gorm:"column:first_name"`
	LastName      *string   `gorm:"column:last_name"`
	DNI           string    `gorm:"column:dni"`
}

func GetApplications(page, pageSize int, isDeleted bool, adminRoleName string) (*dto.ApplicationsListResponse, error) {
	db := config.DB
	adminRolePrefix := strings.TrimSpace(adminRoleName)

	var total int64
	if err := db.Model(&models.Application{}).
		Where("is_deleted = ?", isDeleted).
		Count(&total).Error; err != nil {
		return nil, err
	}

	var apps []models.Application
	if err := db.
		Where("is_deleted = ?", isDeleted).
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&apps).Error; err != nil {
		return nil, err
	}

	if len(apps) == 0 {
		return &dto.ApplicationsListResponse{
			Data:     []dto.ApplicationDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	appIDs := make([]uuid.UUID, 0, len(apps))
	for _, a := range apps {
		appIDs = append(appIDs, a.ID)
	}

	var counts []CountRow
	if err := db.
		Table("user_application_roles").
		Select("application_id, COUNT(DISTINCT user_id) AS users_count").
		Where("is_deleted = FALSE AND revoked_at IS NULL AND application_id IN ?", appIDs).
		Group("application_id").
		Scan(&counts).Error; err != nil {
		return nil, err
	}

	usersCountByApp := make(map[uuid.UUID]int64, len(counts))
	for _, r := range counts {
		usersCountByApp[r.ApplicationID] = r.UsersCount
	}

	var admins []AdminRow
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
		return nil, err
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

	return &dto.ApplicationsListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
