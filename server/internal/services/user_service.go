package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/server/internal/models"
	"github.com/t-saturn/central-user-manager/server/internal/repositories"
)

// CreateUserRequest representa los datos para crear un usuario
type CreateUserRequest struct {
	Email                string
	PasswordHash         string
	FirstName            *string
	LastName             *string
	Phone                *string
	StructuralPositionID *uuid.UUID
	OrganicUnitID        *uuid.UUID
}

// UpdateUserRequest representa los datos para actualizar un usuario
type UpdateUserRequest struct {
	ID                   uuid.UUID
	Email                *string
	FirstName            *string
	LastName             *string
	Phone                *string
	EmailVerified        *bool
	PhoneVerified        *bool
	TwoFactorEnabled     *bool
	Status               *models.UserStatus
	StructuralPositionID *uuid.UUID
	OrganicUnitID        *uuid.UUID
}

// ListUsersRequest representa los parámetros para listar usuarios
type ListUsersRequest struct {
	Page           int
	PageSize       int
	Search         string
	Status         models.UserStatus
	IncludeDeleted bool
}

// ListUsersResponse representa la respuesta de listado de usuarios
type ListUsersResponse struct {
	Users    []models.User
	Total    int64
	Page     int
	PageSize int
}

// UserService define la interfaz para el servicio de usuarios
type UserService interface {
	// CRUD básico
	CreateUser(req CreateUserRequest) (*models.User, error)
	GetUser(id uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(req UpdateUserRequest) (*models.User, error)
	DeleteUser(id uuid.UUID, deletedBy uuid.UUID) error

	// Listado
	ListUsers(req ListUsersRequest) (*ListUsersResponse, error)

	// Operaciones específicas
	UpdatePassword(id uuid.UUID, newPasswordHash string, changedBy uuid.UUID) error
	VerifyEmail(id uuid.UUID) error
	VerifyPhone(id uuid.UUID) error
	EnableTwoFactor(id uuid.UUID) error
	DisableTwoFactor(id uuid.UUID) error
	UpdateLastLogin(id uuid.UUID) error

	// Utilidades
	UserExists(email string) (bool, error)
	GetActiveUsers() ([]models.User, error)
	RestoreUser(id uuid.UUID) error
}

// userService implementa UserService
type userService struct {
	repo repositories.UserRepository
}

// NewUserService crea una nueva instancia del servicio
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser crea un nuevo usuario
func (s *userService) CreateUser(req CreateUserRequest) (*models.User, error) {
	// Validaciones
	if err := s.validateCreateUserRequest(req); err != nil {
		return nil, err
	}

	// Verificar si el email ya existe
	exists, err := s.repo.ExistsByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("error checking email existence: %w", err)
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// Crear el usuario
	user := &models.User{
		Email:                strings.ToLower(strings.TrimSpace(req.Email)),
		PasswordHash:         req.PasswordHash,
		FirstName:            req.FirstName,
		LastName:             req.LastName,
		Phone:                req.Phone,
		StructuralPositionID: req.StructuralPositionID,
		OrganicUnitID:        req.OrganicUnitID,
		Status:               models.UserStatusActive,
		EmailVerified:        false,
		PhoneVerified:        false,
		TwoFactorEnabled:     false,
		IsDeleted:            false,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

// GetUser obtiene un usuario por ID
func (s *userService) GetUser(id uuid.UUID) (*models.User, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

// GetUserByEmail obtiene un usuario por email
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	email = strings.ToLower(strings.TrimSpace(email))
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}

	return user, nil
}

// UpdateUser actualiza un usuario existente
func (s *userService) UpdateUser(req UpdateUserRequest) (*models.User, error) {
	// Validar ID
	if req.ID == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}

	// Obtener usuario existente
	user, err := s.repo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	// Actualizar campos si se proporcionaron
	if req.Email != nil {
		email := strings.ToLower(strings.TrimSpace(*req.Email))
		if email == "" {
			return nil, errors.New("email cannot be empty")
		}

		// Verificar si el nuevo email ya existe (si cambió)
		if user.Email != email {
			exists, err := s.repo.ExistsByEmail(email)
			if err != nil {
				return nil, fmt.Errorf("error checking email existence: %w", err)
			}
			if exists {
				return nil, errors.New("email already exists")
			}
		}
		user.Email = email
	}

	if req.FirstName != nil {
		user.FirstName = req.FirstName
	}

	if req.LastName != nil {
		user.LastName = req.LastName
	}

	if req.Phone != nil {
		user.Phone = req.Phone
	}

	if req.EmailVerified != nil {
		user.EmailVerified = *req.EmailVerified
	}

	if req.PhoneVerified != nil {
		user.PhoneVerified = *req.PhoneVerified
	}

	if req.TwoFactorEnabled != nil {
		user.TwoFactorEnabled = *req.TwoFactorEnabled
	}

	if req.Status != nil {
		user.Status = *req.Status
	}

	if req.StructuralPositionID != nil {
		user.StructuralPositionID = req.StructuralPositionID
	}

	if req.OrganicUnitID != nil {
		user.OrganicUnitID = req.OrganicUnitID
	}

	// Guardar cambios
	if err := s.repo.Update(user); err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return user, nil
}

// DeleteUser elimina un usuario (borrado lógico)
func (s *userService) DeleteUser(id uuid.UUID, deletedBy uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if deletedBy == uuid.Nil {
		return errors.New("invalid deleted_by user ID")
	}

	if err := s.repo.Delete(id, deletedBy); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}

// ListUsers lista usuarios con filtros y paginación
func (s *userService) ListUsers(req ListUsersRequest) (*ListUsersResponse, error) {
	// Validar parámetros
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}

	users, total, err := s.repo.List(req.Page, req.PageSize, req.Search, req.Status, req.IncludeDeleted)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	return &ListUsersResponse{
		Users:    users,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// UpdatePassword actualiza la contraseña de un usuario
func (s *userService) UpdatePassword(id uuid.UUID, newPasswordHash string, changedBy uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if newPasswordHash == "" {
		return errors.New("password hash is required")
	}

	if changedBy == uuid.Nil {
		return errors.New("invalid changed_by user ID")
	}

	if err := s.repo.UpdatePassword(id, newPasswordHash); err != nil {
		return fmt.Errorf("error updating password: %w", err)
	}

	return nil
}

// VerifyEmail marca el email como verificado
func (s *userService) VerifyEmail(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if err := s.repo.VerifyEmail(id); err != nil {
		return fmt.Errorf("error verifying email: %w", err)
	}

	return nil
}

// VerifyPhone marca el teléfono como verificado
func (s *userService) VerifyPhone(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if err := s.repo.VerifyPhone(id); err != nil {
		return fmt.Errorf("error verifying phone: %w", err)
	}

	return nil
}

// EnableTwoFactor habilita la autenticación de dos factores
func (s *userService) EnableTwoFactor(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if err := s.repo.EnableTwoFactor(id); err != nil {
		return fmt.Errorf("error enabling two factor: %w", err)
	}

	return nil
}

// DisableTwoFactor deshabilita la autenticación de dos factores
func (s *userService) DisableTwoFactor(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if err := s.repo.DisableTwoFactor(id); err != nil {
		return fmt.Errorf("error disabling two factor: %w", err)
	}

	return nil
}

// UpdateLastLogin actualiza el timestamp del último login
func (s *userService) UpdateLastLogin(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if err := s.repo.UpdateLastLogin(id); err != nil {
		return fmt.Errorf("error updating last login: %w", err)
	}

	return nil
}

// UserExists verifica si existe un usuario con el email dado
func (s *userService) UserExists(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is required")
	}

	email = strings.ToLower(strings.TrimSpace(email))
	exists, err := s.repo.ExistsByEmail(email)
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %w", err)
	}

	return exists, nil
}

// GetActiveUsers obtiene todos los usuarios activos
func (s *userService) GetActiveUsers() ([]models.User, error) {
	users, err := s.repo.GetActiveUsers()
	if err != nil {
		return nil, fmt.Errorf("error getting active users: %w", err)
	}

	return users, nil
}

// RestoreUser restaura un usuario eliminado
func (s *userService) RestoreUser(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if err := s.repo.RestoreUser(id); err != nil {
		return fmt.Errorf("error restoring user: %w", err)
	}

	return nil
}

// validateCreateUserRequest valida los datos para crear un usuario
func (s *userService) validateCreateUserRequest(req CreateUserRequest) error {
	if req.Email == "" {
		return errors.New("email is required")
	}

	if req.PasswordHash == "" {
		return errors.New("password hash is required")
	}

	// Validar formato de email básico
	email := strings.ToLower(strings.TrimSpace(req.Email))
	if !strings.Contains(email, "@") || len(email) < 5 {
		return errors.New("invalid email format")
	}

	return nil
}
