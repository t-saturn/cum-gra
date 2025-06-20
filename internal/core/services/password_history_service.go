package services

import (
	"fmt"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/shared/dto"
	"github.com/google/uuid"
)

type PasswordHistoryService struct {
	repo repositories.PasswordHistoryRepository
}

func NewPasswordHistoryService(r repositories.PasswordHistoryRepository) *PasswordHistoryService {
	return &PasswordHistoryService{repo: r}
}

func (s *PasswordHistoryService) Create(input dto.CreatePasswordHistoryDTO) error {
	userID := uuid.MustParse(input.UserID)
	exists, err := s.repo.ExistsByUserAndHash(userID, input.PreviousPasswordHash)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("ya existe ese hash de contraseña previa para este usuario")
	}

	history := &domain.PasswordHistory{
		ID:                   uuid.New(),
		UserID:               userID,
		PreviousPasswordHash: input.PreviousPasswordHash,
		ChangedAt:            input.ChangedAt,
		// ChangedBy:            uuid.MustParse(input.ChangedBy),
		IsDeleted: false,
	}

	if input.ChangedBy != "" {
		parsed := uuid.MustParse(input.ChangedBy)
		history.ChangedBy = &parsed
	} else {
		history.ChangedBy = nil
	}

	return s.repo.Create(history)
}

func (s *PasswordHistoryService) GetAll() ([]domain.PasswordHistory, error) {
	return s.repo.GetAllWithUser()
}

func (s *PasswordHistoryService) GetByID(id uuid.UUID) (*domain.PasswordHistory, error) {
	return s.repo.GetByID(id)
}

func (s *PasswordHistoryService) Update(id uuid.UUID, input dto.UpdatePasswordHistoryDTO) error {
	h, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if input.ChangedBy != "" {
		parsed := uuid.MustParse(input.ChangedBy)
		h.ChangedBy = &parsed
	} else {
		h.ChangedBy = nil
	}

	h.PreviousPasswordHash = input.PreviousPasswordHash
	h.ChangedAt = input.ChangedAt
	// h.ChangedBy = uuid.MustParse(input.ChangedBy)

	return s.repo.Update(h)
}

func (s *PasswordHistoryService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
