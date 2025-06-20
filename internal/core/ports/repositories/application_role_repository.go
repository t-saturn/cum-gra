package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type ApplicationRoleRepository interface {
	Create(role *domain.ApplicationRole) error
	GetAll() ([]domain.ApplicationRole, error)
	GetByID(id uuid.UUID) (*domain.ApplicationRole, error)
	Update(role *domain.ApplicationRole) error
	Delete(id uuid.UUID) error
}
