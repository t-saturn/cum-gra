package services

import (
	"fmt"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"
	"server/pkg/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetModules(page, pageSize int, isDeleted bool) (dto.ModulesListResponse, error) {
	db := config.DB

	var total int64
	totalQ := db.Model(&models.Module{})
	if isDeleted {
		totalQ = totalQ.Where("deleted_at IS NOT NULL")
	} else {
		totalQ = totalQ.Where("deleted_at IS NULL")
	}
	if err := totalQ.Count(&total).Error; err != nil {
		logger.Log.Error("Error contando módulos:", err)
		return dto.ModulesListResponse{}, fmt.Errorf("error contando módulos: %w", err)
	}

	var modules []models.Module
	dataQ := db.
		Preload("Application").
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if isDeleted {
		dataQ = dataQ.Where("deleted_at IS NOT NULL")
	} else {
		dataQ = dataQ.Where("deleted_at IS NULL")
	}

	if err := dataQ.Find(&modules).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.ModulesListResponse{
				Data:     []dto.ModuleWithAppDTO{},
				Total:    0,
				Page:     page,
				PageSize: pageSize,
			}, nil
		}
		logger.Log.Error("Error obteniendo módulos:", err)
		return dto.ModulesListResponse{}, fmt.Errorf("error obteniendo módulos: %w", err)
	}

	if len(modules) == 0 {
		return dto.ModulesListResponse{
			Data:     []dto.ModuleWithAppDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	moduleIDs := make([]uuid.UUID, 0, len(modules))
	for _, m := range modules {
		moduleIDs = append(moduleIDs, m.ID)
	}

	type countRow struct {
		ModuleID   uuid.UUID `gorm:"column:module_id"`
		UsersCount int64     `gorm:"column:users_count"`
	}
	var counts []countRow

	if err := db.
		Table("module_role_permissions mrp").
		Select("mrp.module_id, COUNT(DISTINCT uar.user_id) AS users_count").
		Joins(`JOIN user_application_roles uar
					ON uar.application_role_id = mrp.application_role_id
					AND uar.is_deleted = FALSE
					AND uar.revoked_at IS NULL`).
		Joins(`JOIN users u
					ON u.id = uar.user_id
					AND u.is_deleted = FALSE`).
		Where("mrp.is_deleted = FALSE AND mrp.module_id IN ?", moduleIDs).
		Group("mrp.module_id").
		Scan(&counts).Error; err != nil {

		logger.Log.Error("Error contando usuarios por módulo:", err)
		return dto.ModulesListResponse{}, fmt.Errorf("error contando usuarios por módulo: %w", err)
	}

	usersCountByModule := make(map[uuid.UUID]int64, len(counts))
	for _, r := range counts {
		usersCountByModule[r.ModuleID] = r.UsersCount
	}

	out := make([]dto.ModuleWithAppDTO, 0, len(modules))
	for _, m := range modules {
		usersCount := usersCountByModule[m.ID]
		out = append(out, mapper.ToModuleWithAppDTO(m, usersCount))
	}

	return dto.ModulesListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
