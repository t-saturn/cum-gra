package postgres

import (
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/google/uuid"
)

type userModuleRestrictionRepository struct{}

func NewUserModuleRestrictionRepository() repositories.UserModuleRestrictionRepository {
	return &userModuleRestrictionRepository{}
}

func (r *userModuleRestrictionRepository) Create(obj *domain.UserModuleRestriction) error {
	return database.DB.Create(obj).Error
}

func (r *userModuleRestrictionRepository) GetAllWithRelations() ([]domain.UserModuleRestriction, error) {
	var list []domain.UserModuleRestriction
	err := database.DB.
		Preload("User").
		Preload("Module").
		Preload("Application").
		Where("is_deleted = false").
		Find(&list).Error

	return list, err
}

func (r *userModuleRestrictionRepository) GetByID(id uuid.UUID) (*domain.UserModuleRestriction, error) {
	var item domain.UserModuleRestriction
	err := database.DB.
		Preload("User").
		Preload("Module").
		Preload("Application").
		First(&item, "id = ? AND is_deleted = false", id).Error

	return &item, err
}

func (r *userModuleRestrictionRepository) Update(obj *domain.UserModuleRestriction) error {
	return database.DB.Save(obj).Error
}

func (r *userModuleRestrictionRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.UserModuleRestriction{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}

func (r *userModuleRestrictionRepository) ExistsByUserModuleApplication(userID, moduleID, applicationID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.UserModuleRestriction{}).
		Where("user_id = ? AND module_id = ? AND application_id = ? AND is_deleted = false", userID, moduleID, applicationID).
		Count(&count).Error

	return count > 0, err
}
