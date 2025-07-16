package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type ApplicationService struct {
	repo repositories.ApplicationRepository
}

func NewApplicationService(repo repositories.ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		repo: repo,
	}
}

func (s *ApplicationService) IsnameTaken(name string) (bool, error) {
	return s.repo.ExistsByName(name)
}
func (s *ApplicationService) IsNameTakenExceptID(name string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistsByNameExceptID(name, excludeID)
}

func (s *ApplicationService) Create(ctx context.Context, input *dto.CreateApplicationDTO) error {
	entity := &domain.Application{
		Name:         input.Name,
		ClientID:     input.ClientID,
		ClientSecret: input.ClientSecret,
		Domain:       input.Domain,
		Logo:         input.Logo,
		Description:  input.Description,
		CallbackUrls: input.CallbackUrls,
		IsFirstParty: *input.IsFirstParty,
	}

	_, err := s.repo.Create(ctx, entity)
	return err
}

func (s *ApplicationService) GetByID(ctx context.Context, id uuid.UUID) (*dto.ApplicationResponseDTO, error) {
	entity, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.ApplicationResponseDTO{
		ID:           entity.ID,
		Name:         entity.Name,
		ClientID:     entity.ClientID,
		Domain:       entity.Domain,
		Logo:         entity.Logo,
		Description:  entity.Description,
		CallbackUrls: entity.CallbackUrls,
		IsFirstParty: entity.IsFirstParty,
		Status:       entity.Status,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}, nil
}

func (services *ApplicationService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateApplicationDTO) error {
	entity := &domain.Application{}

	if input.Name != nil {
		entity.Name = *input.Name
	}
	if input.ClientID != nil {
		entity.ClientID = *input.ClientID
	}
	if input.ClientSecret != nil {
		entity.ClientSecret = *input.ClientSecret
	}
	if input.Domain != nil {
		entity.Domain = *input.Domain
	}
	if input.Logo != nil {
		entity.Logo = input.Logo
	}
	if input.Description != nil {
		entity.Description = input.Description
	}
	if input.CallbackUrls != nil {
		entity.CallbackUrls = input.CallbackUrls
	}
	if input.IsFirstParty != nil {
		entity.IsFirstParty = *input.IsFirstParty
	}
	if input.Status != nil {
		entity.Status = *input.Status
	}

	_, err := services.repo.Update(ctx, id, entity)
	return err
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
