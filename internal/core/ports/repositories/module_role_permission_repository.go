package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type ModulePermissionRepository interface {
	Create(mp *domain.ModuleRolePermission) error
	GetAllWithRelations() ([]domain.ModuleRolePermission, error)
	GetByID(id uuid.UUID) (*domain.ModuleRolePermission, error)
	Update(mp *domain.ModuleRolePermission) error
	Delete(id uuid.UUID) error
	ExistsByModuleAndRole(moduleID, roleID uuid.UUID) (bool, error)
}
