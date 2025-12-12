package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/pkg/logger"
)

func GetModulesStats() (dto.ModulesStatsResponse, error) {
	db := config.DB
	var resp dto.ModulesStatsResponse

	if err := db.Model(&models.Module{}).Count(&resp.TotalModules).Error; err != nil {
		logger.Log.Error("Error contando total de módulos:", err)
		return resp, err
	}

	if err := db.Model(&models.Module{}).
		Where("status = ? AND deleted_at IS NULL", "active").
		Count(&resp.ActiveModules).Error; err != nil {
		logger.Log.Error("Error contando módulos activos:", err)
		return resp, err
	}

	if err := db.Model(&models.Module{}).
		Where("deleted_at IS NOT NULL").
		Count(&resp.DeletedModules).Error; err != nil {
		logger.Log.Error("Error contando módulos eliminados:", err)
		return resp, err
	}

	if err := db.
		Table("module_role_permissions mrp").
		Joins(`JOIN modules m ON m.id = mrp.module_id AND m.deleted_at IS NULL`).
		Joins(`JOIN user_application_roles uar
					ON uar.application_role_id = mrp.application_role_id
					AND uar.is_deleted = FALSE
					AND uar.revoked_at IS NULL`).
		Joins(`JOIN users u
					ON u.id = uar.user_id
					AND u.is_deleted = FALSE`).
		Where("mrp.is_deleted = FALSE").
		Select("COUNT(DISTINCT u.id)").
		Scan(&resp.TotalUsers).Error; err != nil {
		logger.Log.Error("Error contando usuarios únicos con acceso a módulos:", err)
		return resp, err
	}

	return resp, nil
}
