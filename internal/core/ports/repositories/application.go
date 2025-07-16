package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type ApplicationRepository interface {
	Create(ctx context.Context, sp *domain.Application) (*domain.Application, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Application, error)
	Update(ctx context.Context, id uuid.UUID, sp *domain.Application) (*domain.Application, error)

	ExistsByName(name string) (bool, error)
	ExistsByNameExceptID(name string, excludeID uuid.UUID) (bool, error)
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type ApplicationRoleRepository interface {
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type UserApplicationRoleRepository interface {
}
