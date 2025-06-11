package services

import (
	"github.com/t-saturn/central-user-manager/server/internal/core/domain/entities"
	"github.com/t-saturn/central-user-manager/server/internal/core/ports/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *entities.User) error {
	return s.repo.Create(user)
}
