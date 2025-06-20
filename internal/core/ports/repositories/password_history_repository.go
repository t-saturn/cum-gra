package repositories

import (
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type PasswordHistoryRepository interface {
	Create(p *domain.PasswordHistory) error
	GetAllWithUser() ([]domain.PasswordHistory, error)
	GetByID(id uuid.UUID) (*domain.PasswordHistory, error)
	Update(p *domain.PasswordHistory) error
	Delete(id uuid.UUID) error
	ExistsByUserAndHash(userID uuid.UUID, hash string) (bool, error)
}
