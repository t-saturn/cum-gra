// internal/adapters/repositories/postgres/password_history_repository.go
package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type passwordHistoryRepository struct{}

func NewPasswordHistoryRepository() repositories.PasswordHistoryRepository {
	return &passwordHistoryRepository{}
}

func (r *passwordHistoryRepository) Create(p *domain.PasswordHistory) error {
	return database.DB.Create(p).Error
}

func (r *passwordHistoryRepository) GetAllWithUser() ([]domain.PasswordHistory, error) {
	var items []domain.PasswordHistory
	err := database.DB.
		Preload("User").
		Where("is_deleted = false").
		Find(&items).Error

	return items, err
}

func (r *passwordHistoryRepository) GetByID(id uuid.UUID) (*domain.PasswordHistory, error) {
	var item domain.PasswordHistory
	err := database.DB.
		Preload("User").
		First(&item, "id = ? AND is_deleted = false", id).Error

	return &item, err
}

func (r *passwordHistoryRepository) Update(p *domain.PasswordHistory) error {
	return database.DB.Save(p).Error
}

func (r *passwordHistoryRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.PasswordHistory{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}

func (r *passwordHistoryRepository) ExistsByUserAndHash(userID uuid.UUID, hash string) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.PasswordHistory{}).
		Where("user_id = ? AND previous_password_hash = ? AND is_deleted = false", userID, hash).
		Count(&count).Error

	return count > 0, err
}
