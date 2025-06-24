package services

import (
	"time"

	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"

	"github.com/google/uuid"
)

type ModuleService struct {
	repo repositories.ModuleRepository
}

func NewModuleService(r repositories.ModuleRepository) *ModuleService {
	return &ModuleService{repo: r}
}

func (s *ModuleService) Create(input dto.CreateModuleDTO) error {
	mod := &domain.Module{
		ID:         uuid.New(),
		Item:       &input.Item,
		Name:       input.Name,
		Route:      &input.Route,
		Icon:       &input.Icon,
		SortOrder:  input.SortOrder,
		IsMenuItem: input.IsMenuItem,
		Status:     input.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if input.ApplicationID != "" {
		parsed := uuid.MustParse(input.ApplicationID)
		mod.ApplicationID = &parsed
	} else {
		mod.ApplicationID = nil
	}

	if input.ParentID != "" {
		parentUUID, err := uuid.Parse(input.ParentID)
		if err == nil {
			mod.ParentID = &parentUUID
		}
	}

	return s.repo.Create(mod)
}

func (s *ModuleService) GetAll() ([]domain.Module, error) {
	return s.repo.GetAll()
}

func (s *ModuleService) GetByID(id uuid.UUID) (*domain.Module, error) {
	return s.repo.GetByID(id)
}

func (s *ModuleService) Update(id uuid.UUID, input dto.UpdateModuleDTO) error {
	mod, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if input.ApplicationID != "" {
		parsed := uuid.MustParse(input.ApplicationID)
		mod.ApplicationID = &parsed
	} else {
		mod.ApplicationID = nil
	}

	mod.Item = &input.Item
	mod.Name = input.Name
	mod.Route = &input.Route
	mod.Icon = &input.Icon
	mod.SortOrder = input.SortOrder
	mod.IsMenuItem = input.IsMenuItem
	mod.Status = input.Status
	mod.UpdatedAt = time.Now()

	if input.ParentID != "" {
		if parentUUID, err := uuid.Parse(input.ParentID); err == nil {
			mod.ParentID = &parentUUID
		}
	}

	return s.repo.Update(mod)
}

func (s *ModuleService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
