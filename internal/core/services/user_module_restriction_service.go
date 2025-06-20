// internal/core/services/user_module_restriction_service.go
package services

import (
	"errors"
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/shared/dto"
	"github.com/google/uuid"
)

type UserModuleRestrictionService struct {
	repo repositories.UserModuleRestrictionRepository
}

func NewUserModuleRestrictionService(r repositories.UserModuleRestrictionRepository) *UserModuleRestrictionService {
	return &UserModuleRestrictionService{repo: r}
}

func (s *UserModuleRestrictionService) Create(input dto.CreateUserModuleRestrictionDTO) error {
	// Condición: si restriction_type es limit_permission, max_permission_level es obligatorio
	if input.RestrictionType == "limit_permission" && input.MaxPermissionLevel == "" {
		return errors.New("max_permission_level es obligatorio cuando restriction_type es limit_permission")
	}

	userID := uuid.MustParse(input.UserID)
	moduleID := uuid.MustParse(input.ModuleID)
	appID := uuid.MustParse(input.ApplicationID)

	exists, err := s.repo.ExistsByUserModuleApplication(userID, moduleID, appID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ya existe una restricción para ese usuario, módulo y aplicación")
	}

	restriction := &domain.UserModuleRestriction{
		ID:                 uuid.New(),
		UserID:             userID,
		ModuleID:           moduleID,
		ApplicationID:      appID,
		RestrictionType:    input.RestrictionType,
		MaxPermissionLevel: &input.MaxPermissionLevel,
		/* RestrictionType:    domain.RestrictionType(input.RestrictionType),
		MaxPermissionLevel: domain.PermissionType(input.MaxPermissionLevel), */
		Reason:    &input.Reason,
		ExpiresAt: input.ExpiresAt,
		CreatedAt: time.Now(),
		CreatedBy: uuid.MustParse(input.CreatedBy),
		UpdatedAt: time.Now(),
		//UpdatedBy: uuid.MustParse(input.UpdatedBy),
		IsDeleted: false,
	}

	if input.UpdatedBy != "" {
		parsed := uuid.MustParse(input.UpdatedBy)
		restriction.UpdatedBy = &parsed
	} else {
		restriction.UpdatedBy = nil
	}

	return s.repo.Create(restriction)
}

func (s *UserModuleRestrictionService) GetAll() ([]domain.UserModuleRestriction, error) {
	return s.repo.GetAllWithRelations()
}

func (s *UserModuleRestrictionService) GetByID(id uuid.UUID) (*domain.UserModuleRestriction, error) {
	return s.repo.GetByID(id)
}

func (s *UserModuleRestrictionService) Update(id uuid.UUID, input dto.UpdateUserModuleRestrictionDTO) error {
	r, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if input.RestrictionType == "limit_permission" && input.MaxPermissionLevel == "" {
		return errors.New("max_permission_level es obligatorio cuando restriction_type es limit_permission")
	}

	if input.UpdatedBy != "" {
		parsed := uuid.MustParse(input.UpdatedBy)
		r.UpdatedBy = &parsed
	} else {
		r.UpdatedBy = nil
	}

	r.RestrictionType = input.RestrictionType
	r.MaxPermissionLevel = &input.MaxPermissionLevel
	r.Reason = &input.Reason
	r.ExpiresAt = input.ExpiresAt
	r.UpdatedAt = time.Now()
	// r.UpdatedBy = uuid.MustParse(input.UpdatedBy)

	return s.repo.Update(r)
}

func (s *UserModuleRestrictionService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
