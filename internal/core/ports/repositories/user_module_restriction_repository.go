package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type UserModuleRestrictionRepository interface {
	Create(obj *domain.UserModuleRestriction) error
	GetAllWithRelations() ([]domain.UserModuleRestriction, error)
	GetByID(id uuid.UUID) (*domain.UserModuleRestriction, error)
	Update(obj *domain.UserModuleRestriction) error
	Delete(id uuid.UUID) error
	ExistsByUserModuleApplication(userID, moduleID, applicationID uuid.UUID) (bool, error)
}
