package postgres

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/infrastructure/database"
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
