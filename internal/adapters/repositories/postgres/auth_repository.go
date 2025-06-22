package postgres

import (
	"errors"

	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() *authRepository {
	return &authRepository{db: database.DB}
}

func (r *authRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ? AND is_deleted = false", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("the user does not exist")
	}
	return &user, err
}
