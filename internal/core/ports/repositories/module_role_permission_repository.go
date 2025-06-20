package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type ModulePermissionRepository interface {
	Create(mp *domain.ModuleRolePermission) error
	GetAllWithRelations() ([]domain.ModuleRolePermission, error)
	GetByID(id uuid.UUID) (*domain.ModuleRolePermission, error)
	Update(mp *domain.ModuleRolePermission) error
	Delete(id uuid.UUID) error
	ExistsByModuleAndRole(moduleID, roleID uuid.UUID) (bool, error)
}
