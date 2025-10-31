package services

import (
	"server/internal/config"
	"server/internal/dto"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type roleAssignmentRow struct {
	UserID      uuid.UUID
	UserEmail   string
	UserDNI     string
	UserFName   *string
	UserLName   *string
	AppID       uuid.UUID
	AppName     string
	AppClientID string
	RoleID      uuid.UUID
	RoleName    string
}

func GetRoleAssignments(pageStr, pageSizeStr, isDeletedStr string) (dto.RolesAssigmentsResponseDTO, error) {
	db := config.DB

	page := 1
	pageSize := 20
	isDeleted := false

	if strings.TrimSpace(pageStr) != "" {
		if n, err := strconv.Atoi(pageStr); err == nil && n > 0 {
			page = n
		}
	}
	if strings.TrimSpace(pageSizeStr) != "" {
		if n, err := strconv.Atoi(pageSizeStr); err == nil && n > 0 && n <= 200 {
			pageSize = n
		}
	}
	if strings.TrimSpace(isDeletedStr) != "" {
		if b, err := strconv.ParseBool(isDeletedStr); err == nil {
			isDeleted = b
		}
	}

	buildBase := func(db *gorm.DB) *gorm.DB {
		return db.Table("user_application_roles AS uar").
			Joins("JOIN users u ON u.id = uar.user_id").
			Joins("JOIN applications a ON a.id = uar.application_id").
			Joins("JOIN application_roles r ON r.id = uar.application_role_id").
			Where("uar.is_deleted = ? AND u.is_deleted = ? AND a.is_deleted = ? AND r.is_deleted = ?",
				isDeleted, isDeleted, isDeleted, isDeleted)
	}

	var total int64
	if err := buildBase(db).
		Select("u.id").
		Distinct().
		Count(&total).Error; err != nil {
		return dto.RolesAssigmentsResponseDTO{}, err
	}

	var rows []roleAssignmentRow
	if err := buildBase(db).
		Select(`
			u.id           AS user_id,
			u.email        AS user_email,
			u.dni          AS user_dni,
			u.first_name   AS user_f_name,
			u.last_name    AS user_l_name,
			a.id           AS app_id,
			a.name         AS app_name,
			a.client_id    AS app_client_id,
			r.id           AS role_id,
			r.name         AS role_name
		`).
		Order("COALESCE(u.last_name, '') ASC, COALESCE(u.first_name, '') ASC, a.name ASC, r.name ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Scan(&rows).Error; err != nil {
		return dto.RolesAssigmentsResponseDTO{}, err
	}

	if len(rows) == 0 {
		return dto.RolesAssigmentsResponseDTO{
			Data:     []dto.RoleAssignmentsDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	resultMap := make(map[uuid.UUID]*dto.RoleAssignmentsDTO)
	orderUsers := make([]uuid.UUID, 0, len(rows))

	for _, r := range rows {
		item, ok := resultMap[r.UserID]
		if !ok {
			item = &dto.RoleAssignmentsDTO{
				User: dto.UserMinimalDTO{
					ID:        r.UserID,
					FirstName: r.UserFName,
					LastName:  r.UserLName,
					Email:     r.UserEmail,
					DNI:       r.UserDNI,
				},
				Assignments: []dto.UserAppRoleDTO{},
			}
			resultMap[r.UserID] = item
			orderUsers = append(orderUsers, r.UserID)
		}

		item.Assignments = append(item.Assignments, dto.UserAppRoleDTO{
			Application: dto.AppMinimalDTO{
				ID:       r.AppID,
				Name:     r.AppName,
				ClientID: r.AppClientID,
			},
			Role: dto.RoleMinimalDTO{
				ID:   r.RoleID,
				Name: r.RoleName,
			},
		})
	}

	out := make([]dto.RoleAssignmentsDTO, 0, len(orderUsers))
	for _, uid := range orderUsers {
		out = append(out, *resultMap[uid])
	}

	return dto.RolesAssigmentsResponseDTO{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
