package postgres

import (
	"time"

	"github.com/google/uuid"

	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type moduleRepository struct{}

func NewModuleRepository() repositories.ModuleRepository {
	return &moduleRepository{}
}

func (r *moduleRepository) Create(m *domain.Module) error {
	return database.DB.Create(m).Error
}

func (r *moduleRepository) GetAll() ([]domain.Module, error) {
	var modules []domain.Module
	err := database.DB.Where("is_deleted = false").Find(&modules).Error
	return modules, err
}

func (r *moduleRepository) GetByID(id uuid.UUID) (*domain.Module, error) {
	var module domain.Module
	err := database.DB.First(&module, "id = ? AND is_deleted = false", id).Error
	return &module, err
}

func (r *moduleRepository) Update(m *domain.Module) error {
	return database.DB.Save(m).Error
}

func (r *moduleRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.Module{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}
