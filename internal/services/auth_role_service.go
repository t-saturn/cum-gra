package services

import (
	"time"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/mapper"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

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
		Preload("Module.Children").
		Preload("Module.Parent").
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

	restricted := make(map[uuid.UUID]bool, len(restrictions))
	for _, r := range restrictions {
		restricted[r.ModuleID] = true
	}

	modulesMap := make(map[uuid.UUID]dto.ModuleDTO)
	for _, mp := range modulePerms {
		if mp.Module == nil || restricted[mp.ModuleID] {
			continue
		}
		mod := mapper.ToModuleDTO(mp.Module)
		modulesMap[mod.ID] = mod
	}

	modules := make([]dto.ModuleDTO, 0, len(modulesMap))
	for _, m := range modulesMap {
		modules = append(modules, m)
	}

	return &dto.AuthRoleResponse{
		RoleID:   role.ID.String(),
		RoleName: role.Name,
		Modules:  modules,
	}, nil
}
