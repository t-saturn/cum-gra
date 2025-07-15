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

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
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

func (r *structuralPositionRepository) ExistsByNameExceptID(name string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.StructuralPosition{}).
		Where("LOWER(name) = LOWER(?) AND id <> ? AND is_deleted = false", name, excludeID).
		Count(&count).Error
	return count > 0, err
}

func (r *structuralPositionRepository) ExistsByCodeExceptID(code string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.StructuralPosition{}).
		Where("code = ? AND id <> ? AND is_deleted = false", code, excludeID).
		Count(&count).Error
	return count > 0, err
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
	var existing domain.StructuralPosition
	db := database.DB.WithContext(ctx)

	// Verifica si existe
	if err := db.First(&existing, "id = ? AND is_deleted = false", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// Solo actualiza los campos no nulos
	if sp.Name != "" {
		existing.Name = sp.Name
	}
	if sp.Code != "" {
		existing.Code = sp.Code
	}
	if sp.Level != nil {
		existing.Level = sp.Level
	}
	if sp.Description != nil {
		existing.Description = sp.Description
	}
	if sp.IsActive != existing.IsActive {
		existing.IsActive = sp.IsActive
	}

	existing.UpdatedAt = time.Now()

	if err := db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
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

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type organicUnitRepository struct{}

func NewOrganicUnitRepository() repositories.OrganicUnitRepository {
	return &organicUnitRepository{}
}

func (r *organicUnitRepository) ExistsByName(name string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.OrganicUnit{}).
		Where("name = ? AND is_deleted = false", name).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *organicUnitRepository) ExistsByAcronym(code string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.OrganicUnit{}).
		Where("acronym = ? AND is_deleted = false", code).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *organicUnitRepository) ExistsByID(id uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.OrganicUnit{}).
		Where("id = ? AND is_deleted = false", id).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *organicUnitRepository) Create(ctx context.Context, sp *domain.OrganicUnit) (*domain.OrganicUnit, error) {
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
