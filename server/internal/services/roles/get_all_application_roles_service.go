package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetAllApplicationRoles(applicationID *string) ([]dto.ApplicationRoleSelectDTO, error) {
	db := config.DB

	query := db.Model(&models.ApplicationRole{}).Where("is_deleted = ?", false)

	if applicationID != nil && *applicationID != "" {
		query = query.Where("application_id = ?", *applicationID)
	}

	var roles []models.ApplicationRole
	if err := query.Order("name ASC").Find(&roles).Error; err != nil {
		return nil, err
	}

	result := make([]dto.ApplicationRoleSelectDTO, 0, len(roles))
	for _, r := range roles {
		result = append(result, dto.ApplicationRoleSelectDTO{
			ID:            r.ID.String(),
			Name:          r.Name,
			Description:   r.Description,
			ApplicationID: r.ApplicationID.String(),
		})
	}

	return result, nil
}