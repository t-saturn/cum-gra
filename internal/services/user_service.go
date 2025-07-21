// Package services contiene la lógica de negocio para operaciones relacionadas a usuarios.
package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/config"
	"github.com/t-saturn/central-user-manager/internal/dto"
	"github.com/t-saturn/central-user-manager/internal/models"
	"github.com/t-saturn/central-user-manager/pkg/security"
)

// CreateUser crea un nuevo usuario con los datos proporcionados y lo guarda en la base de datos.
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
