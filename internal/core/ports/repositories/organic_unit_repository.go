package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type OrganicUnitRepository interface {
	Create(unit *domain.OrganicUnit) error
	GetAll() ([]domain.OrganicUnit, error)
	GetByID(id uuid.UUID) (*domain.OrganicUnit, error)
	Update(unit *domain.OrganicUnit) error
	Delete(id uuid.UUID) error
	ExistsByNameOrAcronym(name, acronym string) (bool, error)
}
