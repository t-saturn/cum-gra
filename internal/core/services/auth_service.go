package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/core/ports/services"
)

type AuthService struct {
	repo        repositories.AuthRepository
	hashService services.HashService
}

func NewAuthService(r repositories.AuthRepository, h services.HashService) *AuthService {
	return &AuthService{
		repo:        r,
		hashService: h,
	}
}

func (s *AuthService) Authenticate(email, password string) (uuid.UUID, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return uuid.Nil, err
	}

	if !s.hashService.CheckPasswordHash(password, user.PasswordHash) {
		return uuid.Nil, errors.New("invalid credentials")
	}

	return user.ID, nil
}
