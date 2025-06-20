package repositories

import (
	"github.com/central-user-manager/internal/core/domain"
	"github.com/google/uuid"
)

type PasswordHistoryRepository interface {
	Create(p *domain.PasswordHistory) error
	GetAllWithUser() ([]domain.PasswordHistory, error)
	GetByID(id uuid.UUID) (*domain.PasswordHistory, error)
	Update(p *domain.PasswordHistory) error
	Delete(id uuid.UUID) error
	ExistsByUserAndHash(userID uuid.UUID, hash string) (bool, error)
}
