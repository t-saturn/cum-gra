package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type ModuleRepository interface {
	Create(m *domain.Module) error
	GetAll() ([]domain.Module, error)
	GetByID(id uuid.UUID) (*domain.Module, error)
	Update(m *domain.Module) error
	Delete(id uuid.UUID) error
}
