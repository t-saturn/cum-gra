package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type applicationRoleRepository struct{}

func NewApplicationRoleRepository() repositories.ApplicationRoleRepository {
	return &applicationRoleRepository{}
}

func (r *applicationRoleRepository) Create(role *domain.ApplicationRole) error {
	return database.DB.Create(role).Error
}

func (r *applicationRoleRepository) GetAll() ([]domain.ApplicationRole, error) {
	var roles []domain.ApplicationRole
	err := database.DB.Where("is_deleted = false").Find(&roles).Error
	return roles, err
}

func (r *applicationRoleRepository) GetByID(id uuid.UUID) (*domain.ApplicationRole, error) {
	var role domain.ApplicationRole
	err := database.DB.First(&role, "id = ? AND is_deleted = false", id).Error
	return &role, err
}

func (r *applicationRoleRepository) Update(role *domain.ApplicationRole) error {
	return database.DB.Save(role).Error
}

func (r *applicationRoleRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.ApplicationRole{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}
