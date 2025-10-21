package services

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"

	"github.com/google/uuid"
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

func GetRoleAssignments(page, pageSize int, isDeleted bool) ([]dto.RoleAssignmentsDTO, error) {
	db := config.DB

	var rows []roleAssignmentRow

	q := db.Table("user_application_roles AS uar").
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
		Joins("JOIN users u ON u.id = uar.user_id").
		Joins("JOIN applications a ON a.id = uar.application_id").
		Joins("JOIN application_roles r ON r.id = uar.application_role_id").
		Where("uar.is_deleted = ? AND u.is_deleted = ? AND a.is_deleted = ? AND r.is_deleted = ?",
			isDeleted, isDeleted, isDeleted, isDeleted).
		Order("COALESCE(u.last_name, '') ASC, COALESCE(u.first_name, '') ASC, a.name ASC, r.name ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return []dto.RoleAssignmentsDTO{}, nil
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

	return out, nil
}
