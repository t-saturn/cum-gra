package services

import (
	"fmt"
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/shared/dto"
	"github.com/google/uuid"
)

type StructuralPositionService struct {
	repo repositories.StructuralPositionRepository
}

func NewStructuralPositionService(r repositories.StructuralPositionRepository) *StructuralPositionService {
	return &StructuralPositionService{repo: r}
}

func (s *StructuralPositionService) Create(input dto.CreateStructuralPositionDTO) error {
	// Validar existencia previa
	exists, err := s.repo.ExistsByNameOrCode(input.Name, input.Code)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("ya existe un cargo con ese nombre o código")
	}

	position := &domain.StructuralPosition{
		ID:          uuid.New(),
		Name:        input.Name,
		Code:        input.Code,
		Level:       &input.Level,
		Description: &input.Description,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	return s.repo.Create(position)
}
