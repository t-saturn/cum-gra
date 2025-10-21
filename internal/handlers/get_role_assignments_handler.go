package handlers

import (
	"strconv"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func GetRoleAssignmentsHandler(c fiber.Ctx) error {
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

	// is_deleted (default: false)
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

	// Fila plana para agrupar por usuario
	type row struct {
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

	var rows []row
	if err := q.Scan(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch role assignments",
		})
	}

	// Si no hay filas, devolver wrapper vacío
	if len(rows) == 0 {
		return c.Status(fiber.StatusOK).JSON(dto.RolesAssignmentsResponseDTO{
			Assignments: []dto.RoleAssignmentsDTO{},
		})
	}

	// Agrupar por usuario
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

	// Pasar a slice preservando orden
	out := make([]dto.RoleAssignmentsDTO, 0, len(orderUsers))
	for _, uid := range orderUsers {
		out = append(out, *resultMap[uid])
	}

	return c.Status(fiber.StatusOK).JSON(dto.RolesAssignmentsResponseDTO{
		Assignments: out,
	})
}
