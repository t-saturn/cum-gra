package postgres

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type userApplicationRoleRepository struct{}

func NewUserApplicationRoleRepository() *userApplicationRoleRepository {
	return &userApplicationRoleRepository{}
}

func (r *userApplicationRoleRepository) Create(obj *domain.UserApplicationRole) error {
	return database.DB.Create(obj).Error
}

func (r *userApplicationRoleRepository) GetAll() ([]domain.UserApplicationRole, error) {
	var userApplicationRoles []domain.UserApplicationRole
	err := database.DB.Find(&userApplicationRoles).Error
	if err != nil {
		return nil, err
	}
	return userApplicationRoles, nil
}

func (r *userApplicationRoleRepository) ExistsByUserAndApplication(userID, appID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.UserApplicationRole{}).
		Where("user_id = ? AND application_id = ? AND is_deleted = false", userID, appID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
