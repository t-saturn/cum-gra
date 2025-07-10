package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
)

type StructuralPositionService struct {
	repo repositories.StructuralPositionRepository
}

func NewStructuralPositionService(repo repositories.StructuralPositionRepository) *StructuralPositionService {
	return &StructuralPositionService{
		repo: repo,
	}
}

func (s *StructuralPositionService) Create(ctx context.Context, input *dto.CreateStructuralPositionDTO) (*dto.StructuralPositionResponseDTO, error) {
	entity := &domain.StructuralPosition{
		Name:        input.Name,
		Code:        input.Code,
		Level:       input.Level,
		Description: input.Description,
	}

	created, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &dto.StructuralPositionResponseDTO{
		ID:          created.ID,
		Name:        created.Name,
		Code:        created.Code,
		Level:       created.Level,
		Description: created.Description,
		IsActive:    created.IsActive,
	}, nil
}

func (s *StructuralPositionService) GetByID(ctx context.Context, id uuid.UUID) (*dto.StructuralPositionResponseDTO, error) {
	entity, err := s.repo.GetByID(ctx, id)
	if err != nil || entity == nil {
		return nil, err
	}

	return &dto.StructuralPositionResponseDTO{
		ID:          entity.ID,
		Name:        entity.Name,
		Code:        entity.Code,
		Level:       entity.Level,
		Description: entity.Description,
		IsActive:    entity.IsActive,
	}, nil
}
