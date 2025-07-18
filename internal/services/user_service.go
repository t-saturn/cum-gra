package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/config"
	"github.com/t-saturn/central-user-manager/internal/dto"
	"github.com/t-saturn/central-user-manager/internal/models"
	"github.com/t-saturn/central-user-manager/pkg/security"
)

func CreateUser(input dto.CreateUserDTO) (*models.User, error) {
	passwordHash, err := security.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		Email:        input.Email,
		PasswordHash: passwordHash,
		FirstName:    &input.FirstName,
		LastName:     &input.LastName,
		Phone:        &input.Phone,
		DNI:          input.DNI,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
