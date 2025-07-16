package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	portservices "github.com/t-saturn/central-user-manager/internal/core/ports/services"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type StructuralPositionService struct {
	repo repositories.StructuralPositionRepository
}

func NewStructuralPositionService(repo repositories.StructuralPositionRepository) *StructuralPositionService {
	return &StructuralPositionService{
		repo: repo,
	}
}

func (s *StructuralPositionService) IsNameTaken(name string) (bool, error) {
	return s.repo.ExistsByName(name)
}
func (s *StructuralPositionService) IsCodeTaken(code string) (bool, error) {
	return s.repo.ExistsByCode(code)
}
func (s *StructuralPositionService) IsNameTakenExceptID(name string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistsByNameExceptID(name, excludeID)
}
func (s *StructuralPositionService) IsCodeTakenExceptID(code string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistsByCodeExceptID(code, excludeID)
}

func (s *StructuralPositionService) Create(ctx context.Context, input *dto.CreateStructuralPositionDTO) error {
	entity := &domain.StructuralPosition{
		Name:        input.Name,
		Code:        input.Code,
		Level:       input.Level,
		Description: input.Description,
	}

	_, err := s.repo.Create(ctx, entity)
	return err
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

func (s *StructuralPositionService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateStructuralPositionDTO) error {
	entity := &domain.StructuralPosition{}

	if input.Name != nil {
		entity.Name = *input.Name
	}
	if input.Code != nil {
		entity.Code = *input.Code
	}
	if input.Level != nil {
		entity.Level = input.Level
	}
	if input.Description != nil {
		entity.Description = input.Description
	}
	if input.IsActive != nil {
		entity.IsActive = *input.IsActive
	}

	_, err := s.repo.Update(ctx, id, entity)
	return err
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type OrganicUnitService struct {
	repo repositories.OrganicUnitRepository
}

func NewOrganicUnitService(repo repositories.OrganicUnitRepository) *OrganicUnitService {
	return &OrganicUnitService{
		repo: repo,
	}
}

func (s *OrganicUnitService) IsNameTaken(name string) (bool, error) {
	return s.repo.ExistsByName(name)
}
func (s *OrganicUnitService) IsAcronymTaken(acronym string) (bool, error) {
	return s.repo.ExistsByAcronym(acronym)
}
func (s *OrganicUnitService) IsIdTaken(id uuid.UUID) (bool, error) {
	return s.repo.ExistsByID(id)
}
func (s *OrganicUnitService) IsNameTakenExceptID(name string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistsByNameExceptID(name, excludeID)
}
func (s *OrganicUnitService) IsAcronymTakenExceptID(code string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistsByAcronymExceptID(code, excludeID)
}
func (s *OrganicUnitService) IsIdTakenExeptedID(id string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistsByIDExceptID(id, excludeID)
}

func (s *OrganicUnitService) Create(ctx context.Context, input *dto.CreateOrganicUnitDTO) error {
	entity := &domain.OrganicUnit{
		Name:        input.Name,
		Acronym:     input.Acronym,
		Brand:       input.Brand,
		Description: input.Description,
		ParentID:    input.ParentID,
	}

	_, err := s.repo.Create(ctx, entity)
	return err
}

func (s *OrganicUnitService) GetByID(ctx context.Context, id uuid.UUID) (*dto.OrganicUnitResponseDTO, error) {
	entity, err := s.repo.GetByID(ctx, id)
	if err != nil || entity == nil {
		return nil, err
	}

	return &dto.OrganicUnitResponseDTO{
		ID:          entity.ID,
		Name:        entity.Name,
		Acronym:     entity.Acronym,
		Brand:       entity.Brand,
		Description: entity.Description,
		ParentID:    entity.ParentID,
		IsActive:    entity.IsActive,
	}, nil
}

func (s *OrganicUnitService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateOrganicUnitDTO) error {
	entity := &domain.OrganicUnit{}

	if input.Name != nil {
		entity.Name = *input.Name
	}
	if input.Acronym != nil {
		entity.Acronym = *input.Acronym
	}
	if input.Brand != nil {
		entity.Brand = input.Brand
	}
	if input.Description != nil {
		entity.Description = input.Description
	}
	if input.ParentID != nil {
		entity.ParentID = input.ParentID
	}
	if input.IsActive != nil {
		entity.IsActive = *input.IsActive
	}

	_, err := s.repo.Update(ctx, id, entity)
	return err
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type UserService struct {
	repo        repositories.UserRepository
	hashService portservices.HashService // <-- aquí
}

// nuevo constructor
func NewUserService(repo repositories.UserRepository, hashService portservices.HashService) *UserService {
	return &UserService{
		repo:        repo,
		hashService: hashService,
	}
}

func (s *UserService) IsEmailTaken(email string) (bool, error) {
	return s.repo.ExistByEmail(email)
}

func (s *UserService) IsPhoneTaken(phone string) (bool, error) {
	return s.repo.ExistByPhone(phone)
}

func (s *UserService) IsDniTaken(dni string) (bool, error) {
	return s.repo.ExistByDni(dni)
}

func (s *UserService) IsEmailTakenExceptID(email string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistByEmailExceptID(email, excludeID)
}

func (s *UserService) IsPhoneTakenExceptID(phone string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistByPhoneExceptID(phone, excludeID)
}

func (s *UserService) IsDniTakenExceptID(dni string, excludeID uuid.UUID) (bool, error) {
	return s.repo.ExistByDniExceptID(dni, excludeID)
}

func (s *UserService) IsStructuralPositionIDTaken(id uuid.UUID) (bool, error) {
	return s.repo.StructuralPositionExists(id)
}

func (s *UserService) IsOrganicUnitIDTaken(id uuid.UUID) (bool, error) {
	return s.repo.OrganicUnitExists(id)
}

func (s *UserService) Create(ctx context.Context, input *dto.CreateUserDTO) error {
	hashedPassword, err := s.hashService.HashPassword(input.Password)
	if err != nil {
		return err
	}

	entity := &domain.User{
		Email:                input.Email,
		PasswordHash:         hashedPassword,
		FirstName:            input.FirstName,
		LastName:             input.LastName,
		Phone:                input.Phone,
		DNI:                  input.DNI,
		StructuralPositionID: input.StructuralPositionID,
		OrganicUnitID:        input.OrganicUnitID,
	}

	_, err = s.repo.Create(ctx, entity)
	return err
}
