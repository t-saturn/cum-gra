package services

import (
	"fmt"
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/shared/dto"
	"github.com/google/uuid"
)

type ApplicationService struct {
	repo repositories.ApplicationRepository
}

func NewApplicationService(r repositories.ApplicationRepository) *ApplicationService {
	return &ApplicationService{repo: r}
}

func (s *ApplicationService) Create(input dto.CreateApplicationDTO) error {
	exists, err := s.repo.ExistsByClientID(input.ClientID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("ya existe una aplicación con ese client_id")
	}

	app := &domain.Application{
		ID:           uuid.New(),
		Name:         input.Name,
		ClientID:     input.ClientID,
		ClientSecret: input.ClientSecret,
		Domain:       input.Domain,
		Logo:         &input.Logo,
		Description:  &input.Description,
		CallbackUrls: input.CallbackUrls,
		Scopes:       input.Scopes,
		IsFirstParty: input.IsFirstParty,
		Status:       input.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IsDeleted:    false,
	}

	return s.repo.Create(app)
}

func (s *ApplicationService) GetAll() ([]domain.Application, error) {
	return s.repo.GetAll()
}

func (s *ApplicationService) GetByID(id uuid.UUID) (*domain.Application, error) {
	return s.repo.GetByID(id)
}

func (s *ApplicationService) Update(id uuid.UUID, input dto.UpdateApplicationDTO) error {
	app, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	app.Name = input.Name
	app.ClientSecret = input.ClientSecret
	app.Domain = input.Domain
	app.Logo = &input.Logo
	app.Description = &input.Description
	app.CallbackUrls = input.CallbackUrls
	app.Scopes = input.Scopes
	app.IsFirstParty = input.IsFirstParty
	app.Status = input.Status
	app.UpdatedAt = time.Now()

	return s.repo.Update(app)
}

func (s *ApplicationService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
