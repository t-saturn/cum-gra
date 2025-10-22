package handlers

import (
	"strconv"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func GetUserRestrictionsHandler(c fiber.Ctx) error {
	db := config.DB

	// Paginación
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

	// Filtro de users por is_deleted (default: false)
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

	// 1) Total de usuarios
	var total int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = ?", isDeleted).
		Count(&total).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to count users"})
	}

	// 2) Traer usuarios de la página
	type userRow struct {
		ID        uuid.UUID
		FirstName *string
		LastName  *string
		DNI       string
		Email     string
		Status    string
	}
	var userRows []userRow

	if err := db.Table("users u").
		Select("u.id, u.first_name, u.last_name, u.dni, u.email, u.status").
		Where("u.is_deleted = ?", isDeleted).
		Order("COALESCE(u.last_name, '') ASC, COALESCE(u.first_name, '') ASC, u.email ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Scan(&userRows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch users"})
	}

	// Si no hay usuarios en la página, retorna inmediatamente
	if len(userRows) == 0 {
		return c.Status(fiber.StatusOK).JSON(dto.UsersAppAccessResponse{
			Data:     []dto.UserWithAppsDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
	}

	// 3) IDs de usuarios para traer relaciones
	userIDs := make([]uuid.UUID, 0, len(userRows))
	for _, u := range userRows {
		userIDs = append(userIDs, u.ID)
	}

	// 4) Traer UN rol activo por (user_id, app_id): el más reciente por granted_at
	type uarRow struct {
		UserID      uuid.UUID
		AppID       uuid.UUID
		AppName     string
		AppClientID string
		RoleID      uuid.UUID
		RoleName    string
	}
	var uarRows []uarRow

	// Nota: esta consulta es para PostgreSQL (usa ventana y CTE)
	if err := db.Raw(`
	WITH latest_uar AS (
		SELECT
			uar.user_id,
			uar.application_id,
			uar.application_role_id,
			uar.granted_at,
			ROW_NUMBER() OVER (
				PARTITION BY uar.user_id, uar.application_id
				ORDER BY uar.granted_at DESC
			) AS rn
		FROM user_application_roles uar
		WHERE uar.is_deleted = FALSE
		  AND uar.user_id IN (?)
	)
	SELECT
		l.user_id          AS user_id,
		a.id               AS app_id,
		a.name             AS app_name,
		a.client_id        AS app_client_id,
		ar.id              AS role_id,
		ar.name            AS role_name
	FROM latest_uar l
	JOIN applications a      ON a.id  = l.application_id AND a.is_deleted  = FALSE
	JOIN application_roles ar ON ar.id = l.application_role_id AND ar.is_deleted = FALSE
	WHERE l.rn = 1
`, userIDs).Scan(&uarRows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch user roles"})
	}

	type umrRow struct {
		UserID             uuid.UUID
		AppID              uuid.UUID
		ModuleID           uuid.UUID
		ModuleName         string
		ModuleIcon         *string
		RestrictionType    string
		MaxPermissionLevel *string
	}
	var umrRows []umrRow

	if err := db.Table("user_module_restrictions umr").
		Select(`
			umr.user_id         AS user_id,
			umr.application_id  AS app_id,
			m.id                AS module_id,
			m.name              AS module_name,
			m.icon              AS module_icon,
			umr.restriction_type AS restriction_type,
			umr.max_permission_level AS max_permission_level
		`).
		Joins("JOIN modules m ON m.id = umr.module_id").
		Where("umr.is_deleted = FALSE").
		Where("umr.user_id IN ?", userIDs).
		Scan(&umrRows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch user module restrictions"})
	}

	// 6) Armar estructuras en memoria
	// Map: userID -> item UserWithAppsDTO
	userMap := make(map[uuid.UUID]*dto.UserWithAppsDTO, len(userRows))
	// Map embebido: userID:appID -> *UserAppAccessDTO
	type keyUA struct {
		UserID uuid.UUID
		AppID  uuid.UUID
	}
	uaMap := make(map[keyUA]*dto.UserAppAccessDTO)

	// Inicializar usuarios
	orderUsers := make([]uuid.UUID, 0, len(userRows))
	for _, u := range userRows {
		userMap[u.ID] = &dto.UserWithAppsDTO{
			User: dto.UserBriefDTO{
				ID:        u.ID,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				DNI:       u.DNI,
				Email:     u.Email,
				Status:    u.Status,
			},
			Applications: []dto.UserAppAccessDTO{},
		}
		orderUsers = append(orderUsers, u.ID)
	}

	// Cargar roles por app
	for _, r := range uarRows {
		k := keyUA{UserID: r.UserID, AppID: r.AppID}
		entry, ok := uaMap[k]
		if !ok {
			entry = &dto.UserAppAccessDTO{
				Application: dto.AppMinimalDTO{
					ID:       r.AppID,
					Name:     r.AppName,
					ClientID: r.AppClientID,
				},
				Role: dto.RoleMinimalDTO{
					ID:   r.RoleID,
					Name: r.RoleName,
				},
				RestrictedModules: []dto.RestrictedModuleDTO{},
			}
			uaMap[k] = entry
			// anexar al usuario
			if u := userMap[r.UserID]; u != nil {
				u.Applications = append(u.Applications, *entry)
			}
		} else {
			// Si existiera más de un rol por app/usuario, puedes decidir si:
			// - Reemplazar, o
			// - Ignorar (asumimos unicidad por (user_id, app_id))
			// Aquí NO duplicamos.
		}
	}

	// Cargar módulos restringidos por app
	for _, m := range umrRows {
		k := keyUA{UserID: m.UserID, AppID: m.AppID}
		if entry, ok := uaMap[k]; ok {
			entry.RestrictedModules = append(entry.RestrictedModules, dto.RestrictedModuleDTO{
				ModuleID:           m.ModuleID,
				ModuleName:         m.ModuleName,
				ModuleIcon:         m.ModuleIcon,
				RestrictionType:    m.RestrictionType,
				MaxPermissionLevel: m.MaxPermissionLevel,
			})
		}
	}

	// 7) Pasar a slice respetando el orden de usuarios
	out := make([]dto.UserWithAppsDTO, 0, len(orderUsers))
	for _, uid := range orderUsers {
		out = append(out, *userMap[uid])
	}

	// 8) Responder
	return c.Status(fiber.StatusOK).JSON(dto.UsersAppAccessResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
