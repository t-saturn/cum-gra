package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type StructuralPositionRepository interface {
	Create(position *domain.StructuralPosition) error
	GetAll() ([]domain.StructuralPosition, error)
	GetByID(id uuid.UUID) (*domain.StructuralPosition, error)
	Update(position *domain.StructuralPosition) error
	Delete(id uuid.UUID) error
	ExistsByNameOrCode(name, code string) (bool, error)
}
