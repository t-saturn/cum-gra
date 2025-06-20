package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetAll() ([]domain.User, error)
	GetByID(id uuid.UUID) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
	ExistsByEmailOrDNI(email, dni string) (bool, error)
}
