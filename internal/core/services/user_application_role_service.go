package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
)

type UserApplicationRoleService struct {
	repo repositories.UserApplicationRoleRepository
}

func NewUserApplicationRoleService(r repositories.UserApplicationRoleRepository) *UserApplicationRoleService {
	return &UserApplicationRoleService{repo: r}
}

func (s *UserApplicationRoleService) GetAll() ([]domain.UserApplicationRole, error) {
	return s.repo.GetAll()
}

func (s *UserApplicationRoleService) Create(input dto.CreateUserApplicationRoleDTO) error {
	// Validar si ya existe la relación para ese usuario y esa aplicación
	exists, err := s.repo.ExistsByUserAndApplication(input.UserID, input.ApplicationID)
	if err != nil {
		return fmt.Errorf("error al validar existencia: %w", err)
	}
	if exists {
		return fmt.Errorf("el usuario ya tiene un rol asignado en esta aplicación")
	}

	now := time.Now()
	role := &domain.UserApplicationRole{
		ID:                uuid.New(),
		UserID:            input.UserID,
		ApplicationID:     input.ApplicationID,
		ApplicationRoleID: input.ApplicationRoleID,
		GrantedAt:         now,
		GrantedBy:         input.GrantedBy,
		IsDeleted:         false,
	}

	return s.repo.Create(role)
}
