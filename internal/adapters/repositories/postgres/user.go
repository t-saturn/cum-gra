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

func (r *organicUnitRepository) ExistsByNameExceptID(name string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.OrganicUnit{}).
		Where("LOWER(name) = LOWER(?) AND id <> ? AND is_deleted = false", name, excludeID).
		Count(&count).Error
	return count > 0, err
}

func (r *organicUnitRepository) ExistsByAcronymExceptID(code string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.OrganicUnit{}).
		Where("acronym = ? AND id <> ? AND is_deleted = false", code, excludeID).
		Count(&count).Error
	return count > 0, err
}

func (r *organicUnitRepository) ExistsByIDExceptID(id string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.OrganicUnit{}).
		Where("id = ? AND id <> ? AND is_deleted = false", id, excludeID).
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

func (h *organicUnitRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.OrganicUnit, error) {
	var ou domain.OrganicUnit
	err := database.DB.WithContext(ctx).
		Where("id = ? AND is_deleted = false", id).
		First(&ou).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &ou, nil
}

func (r *organicUnitRepository) Update(ctx context.Context, id uuid.UUID, sp *domain.OrganicUnit) (*domain.OrganicUnit, error) {
	var existing domain.OrganicUnit
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
	if sp.Acronym != "" {
		existing.Acronym = sp.Acronym
	}
	if sp.Brand != nil {
		existing.Brand = sp.Brand
	}
	if sp.Description != nil {
		existing.Description = sp.Description
	}
	if sp.ParentID != nil {
		existing.ParentID = sp.ParentID
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

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type userRepository struct{}

func NewUserRepository() repositories.UserRepository {
	return &userRepository{}
}

func (r *userRepository) ExistByEmail(email string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.User{}).
		Where("email = ? AND is_deleted = false", email).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) ExistByPhone(phone string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.User{}).
		Where("phone = ? AND is_deleted = false", phone).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) ExistByDni(dni string) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.User{}).
		Where("dni = ? AND is_deleted = false", dni).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) ExistByEmailExceptID(email string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.User{}).
		Where("LOWER(email) = LOWER(?) AND id <> ? AND is_deleted = false", email, excludeID).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, err
}

func (r *userRepository) ExistByPhoneExceptID(phone string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.User{}).
		Where("phone = ? AND id <> ? AND is_deleted = false", phone, excludeID).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, err
}

func (r *userRepository) ExistByDniExceptID(dni string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.User{}).
		Where("dni =? AND id <> ? AND is_deleted = false", dni, excludeID).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, err
}
func (r *userRepository) StructuralPositionExists(id uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.StructuralPosition{}).
		Where("id = ? AND is_deleted = false", id).
		Count(&count).Error
	return count > 0, err
}

func (r *userRepository) OrganicUnitExists(id uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.
		Model(&domain.OrganicUnit{}).
		Where("id = ? AND is_deleted = false", id).
		Count(&count).Error
	return count > 0, err
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = "active"
	user.IsDeleted = false

	if err := database.DB.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (h *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var ou domain.User
	err := database.DB.WithContext(ctx).
		Where("id = ? AND is_deleted = false", id).
		First(&ou).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &ou, nil
}

func (r *userRepository) Update(ctx context.Context, id uuid.UUID, user *domain.User) (*domain.User, error) {
	var existing domain.User
	db := database.DB.WithContext(ctx)

	// Verifica si existe
	if err := db.First(&existing, "id = ? AND is_deleted = false", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// Solo actualiza los campos no nulos
	if *user.FirstName != "" {
		existing.FirstName = user.FirstName
	}
	if *user.LastName != "" {
		existing.LastName = user.LastName
	}
	if user.Email != "" {
		existing.Email = user.Email
	}
	if *user.Phone != "" {
		existing.Phone = user.Phone
	}
	if user.DNI != "" {
		existing.DNI = user.DNI
	}
	if user.StructuralPositionID != nil {
		existing.StructuralPositionID = user.StructuralPositionID
	}
	if user.OrganicUnitID != nil {
		existing.OrganicUnitID = user.OrganicUnitID
	}
	if user.PasswordHash != "" {
		existing.PasswordHash = user.PasswordHash
	}

	existing.UpdatedAt = time.Now()

	if err := db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil

}
