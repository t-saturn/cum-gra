package services

import (
	"fmt"
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/core/ports/services"
	"github.com/central-user-manager/internal/shared/dto"

	"github.com/google/uuid"
)

type UserService struct {
	repo        repositories.UserRepository
	hashService services.HashService
}

func NewUserService(r repositories.UserRepository, hs services.HashService) *UserService {
	return &UserService{repo: r, hashService: hs}
}

func (s *UserService) Create(input dto.CreateUserDTO) error {
	exists, err := s.repo.ExistsByEmailOrDNI(input.Email, input.DNI)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("ya existe un usuario con ese email o DNI")
	}

	hashedPassword, err := s.hashService.HashPassword(input.PasswordHash)
	if err != nil {
		return fmt.Errorf("error al encriptar la contraseña")
	}

	user := &domain.User{
		ID:           uuid.New(),
		Email:        input.Email,
		PasswordHash: hashedPassword,
		FirstName:    &input.FirstName,
		LastName:     &input.LastName,
		Phone:        &input.Phone,
		DNI:          input.DNI,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IsDeleted:    false,
	}

	if input.StructuralPositionID != "" {
		parsed, err := uuid.Parse(input.StructuralPositionID)
		if err != nil {
			return fmt.Errorf("invalid StructuralPositionID: %v", err)
		}
		user.StructuralPositionID = &parsed
	} else {
		user.StructuralPositionID = nil
	}
	if input.OrganicUnitID != "" {
		parsed, err := uuid.Parse(input.OrganicUnitID)
		if err != nil {
			return fmt.Errorf("invalid OrganicUnitID: %v", err)
		}
		user.OrganicUnitID = &parsed
	} else {
		user.OrganicUnitID = nil
	}

	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id uuid.UUID) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(id uuid.UUID, input dto.UpdateUserDTO) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if input.StructuralPositionID != "" {
		parsed := uuid.MustParse(input.StructuralPositionID)
		user.StructuralPositionID = &parsed
	} else {
		user.StructuralPositionID = nil
	}
	if input.OrganicUnitID != "" {
		parsed := uuid.MustParse(input.OrganicUnitID)
		user.OrganicUnitID = &parsed
	} else {
		user.OrganicUnitID = nil
	}

	user.FirstName = &input.FirstName
	user.LastName = &input.LastName
	user.Phone = &input.Phone
	user.Status = input.Status
	// user.Status = domain.UserStatus(input.Status)
	user.UpdatedAt = time.Now()

	return s.repo.Update(user)
}

func (s *UserService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
