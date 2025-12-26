package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetAllModules(onlyActive bool, applicationID *string) ([]dto.ModuleSelectDTO, error) {
	db := config.DB

	query := db.Model(&models.Module{}).Where("deleted_at IS NULL")

	if onlyActive {
		query = query.Where("status = ?", "active")
	}

	if applicationID != nil && *applicationID != "" {
		query = query.Where("application_id = ?", *applicationID)
	}

	var modules []models.Module
	if err := query.Order("sort_order ASC, name ASC").Find(&modules).Error; err != nil {
		return nil, err
	}

	result := make([]dto.ModuleSelectDTO, 0, len(modules))
	for _, m := range modules {
		var parentID, appID *string
		if m.ParentID != nil {
			pid := m.ParentID.String()
			parentID = &pid
		}
		if m.ApplicationID != nil {
			aid := m.ApplicationID.String()
			appID = &aid
		}

		result = append(result, dto.ModuleSelectDTO{
			ID:            m.ID.String(),
			Name:          m.Name,
			Item:          m.Item,
			Icon:          m.Icon,
			ParentID:      parentID,
			ApplicationID: appID,
			Status:        m.Status,
		})
	}

	return result, nil
}