package postgres

import (
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/google/uuid"
)

type structuralPositionRepository struct{}

func NewStructuralPositionRepository() repositories.StructuralPositionRepository {
	return &structuralPositionRepository{}
}

func (r *structuralPositionRepository) Create(position *domain.StructuralPosition) error {
	return database.DB.Create(position).Error
}

func (r *structuralPositionRepository) ExistsByNameOrCode(name, code string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.StructuralPosition{}).
		Where("name = ? OR code = ?", name, code).
		Count(&count).Error
	return count > 0, err
}

func (r *structuralPositionRepository) GetAll() ([]domain.StructuralPosition, error) {
	var positions []domain.StructuralPosition
	err := database.DB.Where("is_deleted = false").Find(&positions).Error
	return positions, err
}

func (r *structuralPositionRepository) GetByID(id uuid.UUID) (*domain.StructuralPosition, error) {
	var position domain.StructuralPosition
	err := database.DB.First(&position, "id = ? AND is_deleted = false", id).Error
	return &position, err
}

func (r *structuralPositionRepository) Update(position *domain.StructuralPosition) error {
	return database.DB.Save(position).Error
}

func (r *structuralPositionRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.StructuralPosition{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}
