package repositories

import (
	"github.com/t-saturn/central-user-manager/server/internal/adapters/secondary/persistence/postgres"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/secondary/persistence/postgres/models"
	"github.com/t-saturn/central-user-manager/server/internal/core/domain/entities"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *entities.User) error {
	model := models.UserModel{
		Name:     user.Name,
		LastName: user.LastName,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	return postgres.DB.Create(&model).Error
}
