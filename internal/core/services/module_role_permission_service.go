package services

import (
	"fmt"
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/shared/dto"
	"github.com/google/uuid"
)

type ModulePermissionService struct {
	repo repositories.ModulePermissionRepository
}

func NewModulePermissionService(r repositories.ModulePermissionRepository) *ModulePermissionService {
	return &ModulePermissionService{repo: r}
}

func (s *ModulePermissionService) Create(input dto.CreateModuleRolePermissionDTO) error {
	moduleID := uuid.MustParse(input.ModuleID)
	roleID := uuid.MustParse(input.ApplicationRoleID)

	exists, err := s.repo.ExistsByModuleAndRole(moduleID, roleID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("ya existe un permiso para ese módulo y rol")
	}

	perm := &domain.ModuleRolePermission{
		ID:                uuid.New(),
		ModuleID:          moduleID,
		ApplicationRoleID: roleID,
		PermissionType:    input.PermissionType,
		// PermissionType:    domain.PermissionType(input.PermissionType),
		CreatedAt: time.Now(),
		IsDeleted: false,
	}

	return s.repo.Create(perm)
}

func (s *ModulePermissionService) GetAll() ([]domain.ModuleRolePermission, error) {
	return s.repo.GetAllWithRelations()
}

func (s *ModulePermissionService) GetByID(id uuid.UUID) (*domain.ModuleRolePermission, error) {
	return s.repo.GetByID(id)
}

func (s *ModulePermissionService) Update(id uuid.UUID, input dto.UpdateModuleRolePermissionDTO) error {
	perm, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	perm.PermissionType = input.PermissionType
	return s.repo.Update(perm)
}

func (s *ModulePermissionService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
