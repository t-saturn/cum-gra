package services

import (
	"github.com/t-saturn/central-user-manager/server/internal/models"
	"github.com/t-saturn/central-user-manager/server/internal/repositories"

	"github.com/google/uuid"
)

type ApplicationService interface {
	CreateApplication(input *models.Application) (*models.Application, error)
}

type applicationService struct {
	repo repositories.ApplicationRepository
}

func NewApplicationService(repo repositories.ApplicationRepository) ApplicationService {
	return &applicationService{repo: repo}
}

func (s *applicationService) CreateApplication(input *models.Application) (*models.Application, error) {
	input.ID = uuid.New()
	err := s.repo.Create(input)
	if err != nil {
		return nil, err
	}
	return input, nil
}
