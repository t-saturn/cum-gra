package services

import (
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/pkg/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserRoleAndModules(userID uuid.UUID, clientID string) (*dto.AuthRoleResponse, error) {
	db := config.DB

	var app models.Application
	if err := db.Where("client_id = ?", clientID).First(&app).Error; err != nil {
		return nil, err
	}

	var userAppRole models.UserApplicationRole
	if err := db.
		Preload("ApplicationRole").
		Where("user_id = ? AND application_id = ? AND is_deleted = false", userID, app.ID).
		First(&userAppRole).Error; err != nil {
		return nil, err
	}

	role := userAppRole.ApplicationRole
	if role == nil {
		return nil, gorm.ErrRecordNotFound
	}

	var modulePerms []models.ModuleRolePermission
	if err := db.
		Preload("Module", "deleted_at IS NULL").
		Preload("Module.Children", "deleted_at IS NULL").
		Preload("Module.Parent", "deleted_at IS NULL").
		Where("application_role_id = ? AND is_deleted = false", role.ID).
		Find(&modulePerms).Error; err != nil {
		logger.Log.Error("Error obteniendo módulos del rol:", err)
		return nil, err
	}

	now := time.Now()
	var restrictions []models.UserModuleRestriction
	if err := db.
		Where("user_id = ? AND application_id = ? AND is_deleted = false AND (expires_at IS NULL OR expires_at > ?)", userID, app.ID, now).
		Find(&restrictions).Error; err != nil {
		logger.Log.Error("Error obteniendo restricciones de usuario:", err)
		return nil, err
	}

	// Mapa de restricciones por módulo
	restrictionMap := make(map[uuid.UUID]models.UserModuleRestriction, len(restrictions))
	for _, r := range restrictions {
		restrictionMap[r.ModuleID] = r
	}

	// Construir módulos con permisos y restricciones
	modulesMap := make(map[uuid.UUID]dto.ModuleWithPerms)
	for _, mp := range modulePerms {
		if mp.Module == nil {
			continue
		}

		mod := toModuleWithPerms(mp.Module, mp.PermissionType, restrictionMap)
		modulesMap[mp.ModuleID] = mod
	}

	modules := make([]dto.ModuleWithPerms, 0, len(modulesMap))
	for _, m := range modulesMap {
		modules = append(modules, m)
	}

	return &dto.AuthRoleResponse{
		RoleID:   role.ID.String(),
		RoleName: role.Name,
		Modules:  modules,
	}, nil
}

func toModuleWithPerms(m *models.Module, permType string, restrictions map[uuid.UUID]models.UserModuleRestriction) dto.ModuleWithPerms {
	mod := dto.ModuleWithPerms{
		ID:             m.ID.String(),
		Item:           m.Item,
		Name:           m.Name,
		Route:          m.Route,
		Icon:           m.Icon,
		SortOrder:      m.SortOrder,
		Status:         m.Status,
		PermissionType: permType,
	}

	if m.ParentID != nil {
		parentStr := m.ParentID.String()
		mod.ParentID = &parentStr
	}

	// Agregar restricción si existe
	if r, exists := restrictions[m.ID]; exists {
		mod.Restriction = &dto.ModuleRestriction{
			RestrictionType:    r.RestrictionType,
			MaxPermissionLevel: r.MaxPermissionLevel,
			Reason:             r.Reason,
		}
		if r.ExpiresAt != nil {
			exp := r.ExpiresAt.Format(time.RFC3339)
			mod.Restriction.ExpiresAt = &exp
		}
	}

	// Procesar hijos si existen
	if len(m.Children) > 0 {
		mod.Children = make([]dto.ModuleWithPerms, 0, len(m.Children))
		for _, child := range m.Children {
			// Los hijos heredan el permiso del padre (o podrías buscar su permiso específico)
			childMod := toModuleWithPerms(&child, permType, restrictions)
			mod.Children = append(mod.Children, childMod)
		}
	}

	return mod
}