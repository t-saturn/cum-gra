package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type ModuleRepository interface {
	Create(m *domain.Module) error
	GetAll() ([]domain.Module, error)
	GetByID(id uuid.UUID) (*domain.Module, error)
	Update(m *domain.Module) error
	Delete(id uuid.UUID) error
}
