package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type StructuralPositionRepository interface {
	Create(position *domain.StructuralPosition) error
	GetAll() ([]domain.StructuralPosition, error)
	GetByID(id uuid.UUID) (*domain.StructuralPosition, error)
	Update(position *domain.StructuralPosition) error
	Delete(id uuid.UUID) error
	ExistsByNameOrCode(name, code string) (bool, error)
}
