package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
)

type OrganicUnitService struct {
	repo repositories.OrganicUnitRepository
}

func NewOrganicUnitService(r repositories.OrganicUnitRepository) *OrganicUnitService {
	return &OrganicUnitService{repo: r}
}

func (s *OrganicUnitService) Create(input dto.CreateOrganicUnitDTO) error {
	exists, err := s.repo.ExistsByNameOrAcronym(input.Name, input.Acronym)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("ya existe una unidad con ese nombre o acrónimo")
	}

	unit := &domain.OrganicUnit{
		ID:          uuid.New(),
		Name:        input.Name,
		Acronym:     input.Acronym,
		Brand:       &input.Brand,
		Level:       &input.Level,
		Description: &input.Description,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	if input.ParentID != "" {
		if parentUUID, err := uuid.Parse(input.ParentID); err == nil {
			unit.ParentID = &parentUUID
		}
	}

	return s.repo.Create(unit)
}

func (s *OrganicUnitService) GetAll() ([]domain.OrganicUnit, error) {
	return s.repo.GetAll()
}

func (s *OrganicUnitService) GetByID(id uuid.UUID) (*domain.OrganicUnit, error) {
	return s.repo.GetByID(id)
}

func (s *OrganicUnitService) Update(id uuid.UUID, input dto.UpdateOrganicUnitDTO) error {
	unit, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	unit.Name = input.Name
	unit.Acronym = input.Acronym
	unit.Brand = &input.Brand
	unit.Level = &input.Level
	unit.Description = &input.Description
	unit.UpdatedAt = time.Now()

	if input.ParentID != "" {
		if parentUUID, err := uuid.Parse(input.ParentID); err == nil {
			unit.ParentID = &parentUUID
		}
	}

	return s.repo.Update(unit)
}

func (s *OrganicUnitService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
