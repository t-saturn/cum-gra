package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type ApplicationRepository interface {
	Create(app *domain.Application) error
	GetAll() ([]domain.Application, error)
	GetByID(id uuid.UUID) (*domain.Application, error)
	Update(app *domain.Application) error
	Delete(id uuid.UUID) error
	ExistsByClientID(clientID string) (bool, error)
}
