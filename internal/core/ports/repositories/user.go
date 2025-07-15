package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type StructuralPositionRepository interface {
	Create(ctx context.Context, sp *domain.StructuralPosition) (*domain.StructuralPosition, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.StructuralPosition, error)
	Update(ctx context.Context, id uuid.UUID, sp *domain.StructuralPosition) (*domain.StructuralPosition, error)
	// Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) error

	ExistsByName(name string) (bool, error)
	ExistsByCode(code string) (bool, error)
	ExistsByNameExceptID(name string, excludeID uuid.UUID) (bool, error)
	ExistsByCodeExceptID(code string, excludeID uuid.UUID) (bool, error)
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type OrganicUnitRepository interface {
	Create(ctx context.Context, ou *domain.OrganicUnit) (*domain.OrganicUnit, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.OrganicUnit, error)
	Update(ctx context.Context, id uuid.UUID, ou *domain.OrganicUnit) (*domain.OrganicUnit, error)

	ExistsByName(name string) (bool, error)
	ExistsByAcronym(code string) (bool, error)
	ExistsByID(id uuid.UUID) (bool, error)
	ExistsByNameExceptID(name string, excludeID uuid.UUID) (bool, error)
	ExistsByAcronymExceptID(code string, excludeID uuid.UUID) (bool, error)
	ExistsByIDExceptID(id string, excludeID uuid.UUID) (bool, error)
}
