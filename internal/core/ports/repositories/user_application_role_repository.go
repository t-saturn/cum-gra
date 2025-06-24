package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type UserApplicationRoleRepository interface {
	Create(obj *domain.UserApplicationRole) error
	GetAll() ([]domain.UserApplicationRole, error)
	ExistsByUserAndApplication(userID, appID uuid.UUID) (bool, error)
}
