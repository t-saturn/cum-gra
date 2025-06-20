package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
)

type ApplicationRoleService struct {
	repo repositories.ApplicationRoleRepository
}

func NewApplicationRoleService(r repositories.ApplicationRoleRepository) *ApplicationRoleService {
	return &ApplicationRoleService{repo: r}
}

func (s *ApplicationRoleService) Create(input dto.CreateApplicationRoleDTO) error {
	role := &domain.ApplicationRole{
		ID:            uuid.New(),
		Name:          input.Name,
		Description:   &input.Description,
		ApplicationID: uuid.MustParse(input.ApplicationID),
		Permissions:   input.Permissions,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		IsDeleted:     false,
	}
	return s.repo.Create(role)
}

func (s *ApplicationRoleService) GetAll() ([]domain.ApplicationRole, error) {
	return s.repo.GetAll()
}

func (s *ApplicationRoleService) GetByID(id uuid.UUID) (*domain.ApplicationRole, error) {
	return s.repo.GetByID(id)
}

func (s *ApplicationRoleService) Update(id uuid.UUID, input dto.UpdateApplicationRoleDTO) error {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	role.Name = input.Name
	role.Description = &input.Description
	role.Permissions = input.Permissions
	role.UpdatedAt = time.Now()

	return s.repo.Update(role)
}

func (s *ApplicationRoleService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
