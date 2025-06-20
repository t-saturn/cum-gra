package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type userRepository struct{}

func NewUserRepository() repositories.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *domain.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := database.DB.Where("is_deleted = false").Find(&users).Error
	return users, err
}

func (r *userRepository) GetByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := database.DB.First(&user, "id = ? AND is_deleted = false", id).Error
	return &user, err
}

func (r *userRepository) Update(user *domain.User) error {
	return database.DB.Save(user).Error
}

func (r *userRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}

func (r *userRepository) ExistsByEmailOrDNI(email, dni string) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.User{}).
		Where("email = ? OR dni = ?", email, dni).
		Count(&count).Error
	return count > 0, err
}
