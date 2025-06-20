package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type modulePermissionRepository struct{}

func NewModulePermissionRepository() repositories.ModulePermissionRepository {
	return &modulePermissionRepository{}
}

func (r *modulePermissionRepository) Create(mp *domain.ModuleRolePermission) error {
	return database.DB.Create(mp).Error
}

func (r *modulePermissionRepository) GetAllWithRelations() ([]domain.ModuleRolePermission, error) {
	var list []domain.ModuleRolePermission
	err := database.DB.
		Preload("Module").
		Preload("ApplicationRole").
		Where("is_deleted = false").
		Find(&list).Error
	return list, err
}

func (r *modulePermissionRepository) GetByID(id uuid.UUID) (*domain.ModuleRolePermission, error) {
	var mp domain.ModuleRolePermission
	err := database.DB.
		Preload("Module").
		Preload("ApplicationRole").
		First(&mp, "id = ? AND is_deleted = false", id).Error
	return &mp, err
}

func (r *modulePermissionRepository) Update(mp *domain.ModuleRolePermission) error {
	return database.DB.Save(mp).Error
}

func (r *modulePermissionRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.ModuleRolePermission{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}

func (r *modulePermissionRepository) ExistsByModuleAndRole(moduleID, roleID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.ModuleRolePermission{}).
		Where("module_id = ? AND application_role_id = ? AND is_deleted = false", moduleID, roleID).
		Count(&count).Error
	return count > 0, err
}
