package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type UserModuleRestrictionRepository interface {
	Create(obj *domain.UserModuleRestriction) error
	GetAllWithRelations() ([]domain.UserModuleRestriction, error)
	GetByID(id uuid.UUID) (*domain.UserModuleRestriction, error)
	Update(obj *domain.UserModuleRestriction) error
	Delete(id uuid.UUID) error
	ExistsByUserModuleApplication(userID, moduleID, applicationID uuid.UUID) (bool, error)
}
