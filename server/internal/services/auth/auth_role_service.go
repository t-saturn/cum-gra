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

	// Obtener permisos SIN preload de children
	var modulePerms []models.ModuleRolePermission
	if err := db.
		Preload("Module", "deleted_at IS NULL").
		Where("application_role_id = ? AND is_deleted = false", role.ID).
		Find(&modulePerms).Error; err != nil {
		logger.Log.Error("Error obteniendo módulos del rol:", err)
		return nil, err
	}

	// Obtener restricciones activas del usuario
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

	// Construir lista de módulos SIN duplicados y SIN children automáticos
	modulesMap := make(map[uuid.UUID]dto.ModuleWithPerms)
	for _, mp := range modulePerms {
		if mp.Module == nil {
			continue
		}

		// Si ya existe, saltar (evitar duplicados)
		if _, exists := modulesMap[mp.ModuleID]; exists {
			continue
		}

		mod := dto.ModuleWithPerms{
			ID:             mp.Module.ID.String(),
			Item:           mp.Module.Item,
			Name:           mp.Module.Name,
			Route:          mp.Module.Route,
			Icon:           mp.Module.Icon,
			SortOrder:      mp.Module.SortOrder,
			Status:         mp.Module.Status,
			PermissionType: mp.PermissionType,
			// NO incluir children aquí
		}

		if mp.Module.ParentID != nil {
			parentStr := mp.Module.ParentID.String()
			mod.ParentID = &parentStr
		}

		// Agregar restricción si existe
		if r, exists := restrictionMap[mp.Module.ID]; exists {
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