package services

import (
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
)

func CreateUser(input dto.CreateUserDTO) (*models.User, error) {

	user := &models.User{
		ID:        uuid.New(),
		Email:     input.Email,
		DNI:       input.DNI,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
