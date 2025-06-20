package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type OrganicUnitRepository interface {
	Create(unit *domain.OrganicUnit) error
	GetAll() ([]domain.OrganicUnit, error)
	GetByID(id uuid.UUID) (*domain.OrganicUnit, error)
	Update(unit *domain.OrganicUnit) error
	Delete(id uuid.UUID) error
	ExistsByNameOrAcronym(name, acronym string) (bool, error)
}
