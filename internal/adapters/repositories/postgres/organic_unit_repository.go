package postgres

import (
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/google/uuid"
)

type organicUnitRepository struct{}

func NewOrganicUnitRepository() repositories.OrganicUnitRepository {
	return &organicUnitRepository{}
}

func (r *organicUnitRepository) Create(unit *domain.OrganicUnit) error {
	return database.DB.Create(unit).Error
}

func (r *organicUnitRepository) GetAll() ([]domain.OrganicUnit, error) {
	var units []domain.OrganicUnit
	err := database.DB.Where("is_deleted = false").Find(&units).Error
	return units, err
}

func (r *organicUnitRepository) GetByID(id uuid.UUID) (*domain.OrganicUnit, error) {
	var unit domain.OrganicUnit
	err := database.DB.First(&unit, "id = ? AND is_deleted = false", id).Error
	return &unit, err
}

func (r *organicUnitRepository) Update(unit *domain.OrganicUnit) error {
	return database.DB.Save(unit).Error
}

func (r *organicUnitRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.OrganicUnit{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}

func (r *organicUnitRepository) ExistsByNameOrAcronym(name, acronym string) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.OrganicUnit{}).
		Where("name = ? OR acronym = ?", name, acronym).
		Count(&count).Error
	return count > 0, err
}
