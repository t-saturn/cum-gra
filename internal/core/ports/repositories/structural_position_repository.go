package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
)

type StructuralPositionRepository interface {
	Create(position *domain.StructuralPosition) error
	ExistsByNameOrCode(name, code string) (bool, error)
}
