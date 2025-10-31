package services

import (
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/pkg/security"

	"github.com/google/uuid"
)

func CreateUser(input dto.CreateUserDTO) (*models.User, error) {
	argon := security.NewArgon2Service()
	passwordHash, err := argon.HashPassword(input.Password)

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
