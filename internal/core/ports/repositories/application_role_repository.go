package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type ApplicationRoleRepository interface {
	Create(role *domain.ApplicationRole) error
	GetAll() ([]domain.ApplicationRole, error)
	GetByID(id uuid.UUID) (*domain.ApplicationRole, error)
	Update(role *domain.ApplicationRole) error
	Delete(id uuid.UUID) error
}
