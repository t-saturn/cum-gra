package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type structuralPositionRepository struct{}

func NewStructuralPositionRepository() repositories.StructuralPositionRepository {
	return &structuralPositionRepository{}
}

func (r *structuralPositionRepository) ExistsByName(name string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.StructuralPosition{}).
		Where("name = ? AND is_deleted = false", name).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *structuralPositionRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.StructuralPosition{}).
		Where("code = ? AND is_deleted = false", code).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *structuralPositionRepository) Create(ctx context.Context, sp *domain.StructuralPosition) (*domain.StructuralPosition, error) {
	sp.ID = uuid.New()
	sp.CreatedAt = time.Now()
	sp.UpdatedAt = time.Now()
	sp.IsActive = true
	sp.IsDeleted = false

	if err := database.DB.WithContext(ctx).Create(sp).Error; err != nil {
		return nil, err
	}

	return sp, nil
}

func (r *structuralPositionRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.StructuralPosition, error) {
	var sp domain.StructuralPosition
	err := database.DB.WithContext(ctx).
		Where("id = ? AND is_deleted = false", id).
		First(&sp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &sp, nil
}

func (r *structuralPositionRepository) Update(ctx context.Context, id uuid.UUID, sp *domain.StructuralPosition) (*domain.StructuralPosition, error) {
	sp.UpdatedAt = time.Now()
	err := database.DB.WithContext(ctx).
		Model(&domain.StructuralPosition{}).
		Where("id = ? AND is_deleted = false", id).
		Updates(sp).Error

	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, id)
}

func (r *structuralPositionRepository) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) error {
	now := time.Now()
	return database.DB.WithContext(ctx).
		Model(&domain.StructuralPosition{}).
		Where("id = ? AND is_deleted = false", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": now,
			"deleted_by": deletedBy,
		}).Error
}
