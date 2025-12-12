package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetModuleRolePermissionsStats() (*dto.ModuleRolePermissionsStatsResponse, error) {
	db := config.DB

	var totalPermissions int64
	if err := db.Model(&models.ModuleRolePermission{}).
		Count(&totalPermissions).Error; err != nil {
		return nil, err
	}

	var activePermissions int64
	if err := db.Model(&models.ModuleRolePermission{}).
		Where("is_deleted = FALSE").
		Count(&activePermissions).Error; err != nil {
		return nil, err
	}

	var deletedPermissions int64
	if err := db.Model(&models.ModuleRolePermission{}).
		Where("is_deleted = TRUE").
		Count(&deletedPermissions).Error; err != nil {
		return nil, err
	}

	var uniqueModules int64
	if err := db.Model(&models.ModuleRolePermission{}).
		Where("is_deleted = FALSE").
		Distinct("module_id").
		Count(&uniqueModules).Error; err != nil {
		return nil, err
	}

	var uniqueRoles int64
	if err := db.Model(&models.ModuleRolePermission{}).
		Where("is_deleted = FALSE").
		Distinct("application_role_id").
		Count(&uniqueRoles).Error; err != nil {
		return nil, err
	}

	// Contar por tipo de permiso
	var permissionsByTypeRows []struct {
		PermissionType string `gorm:"column:permission_type"`
		Count          int64  `gorm:"column:count"`
	}
	if err := db.Model(&models.ModuleRolePermission{}).
		Select("permission_type, COUNT(*) as count").
		Where("is_deleted = FALSE").
		Group("permission_type").
		Scan(&permissionsByTypeRows).Error; err != nil {
		return nil, err
	}

	permissionsByType := make(map[string]int64)
	for _, row := range permissionsByTypeRows {
		permissionsByType[row.PermissionType] = row.Count
	}

	return &dto.ModuleRolePermissionsStatsResponse{
		TotalPermissions:   totalPermissions,
		ActivePermissions:  activePermissions,
		DeletedPermissions: deletedPermissions,
		UniqueModules:      uniqueModules,
		UniqueRoles:        uniqueRoles,
		PermissionsByType:  permissionsByType,
	}, nil
}