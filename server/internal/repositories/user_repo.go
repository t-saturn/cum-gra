package repositories

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/server/internal/models"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
	"gorm.io/gorm"
)

// UserRepository define la interfaz para operaciones de usuario
type UserRepository interface {
	// CRUD básico
	Create(user *models.User) error
	GetByID(id uuid.UUID) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID, deletedBy uuid.UUID) error

	// Listado con filtros
	List(page, pageSize int, search string, status models.UserStatus, includeDeleted bool) ([]models.User, int64, error)

	// Operaciones específicas
	UpdatePassword(id uuid.UUID, newPasswordHash string) error
	VerifyEmail(id uuid.UUID) error
	VerifyPhone(id uuid.UUID) error
	EnableTwoFactor(id uuid.UUID) error
	DisableTwoFactor(id uuid.UUID) error
	UpdateLastLogin(id uuid.UUID) error

	// Utilidades
	ExistsByEmail(email string) (bool, error)
	GetActiveUsers() ([]models.User, error)
	RestoreUser(id uuid.UUID) error
}

// userRepository implementa UserRepository
type userRepository struct{}

// NewUserRepository crea una nueva instancia del repositorio
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// Create crea un nuevo usuario
func (r *userRepository) Create(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	// Verificar si el email ya existe
	exists, err := r.ExistsByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("error checking email existence: %w", err)
	}
	if exists {
		return errors.New("email already exists")
	}

	return database.DB.Create(user).Error
}

// GetByID obtiene un usuario por ID
func (r *userRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := database.DB.Where("id = ? AND is_deleted = ?", id, false).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail obtiene un usuario por email
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ? AND is_deleted = ?", email, false).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Update actualiza un usuario existente
func (r *userRepository) Update(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	// Verificar que el usuario existe
	existing, err := r.GetByID(user.ID)
	if err != nil {
		return err
	}

	// Si se cambió el email, verificar que no exista otro usuario con ese email
	if existing.Email != user.Email {
		exists, err := r.ExistsByEmail(user.Email)
		if err != nil {
			return fmt.Errorf("error checking email existence: %w", err)
		}
		if exists {
			return errors.New("email already exists")
		}
	}

	return database.DB.Save(user).Error
}

// Delete realiza borrado lógico de un usuario
func (r *userRepository) Delete(id uuid.UUID, deletedBy uuid.UUID) error {
	user, err := r.GetByID(id)
	if err != nil {
		return err
	}

	user.SoftDelete(deletedBy)
	return database.DB.Save(user).Error
}

// List obtiene usuarios con paginación y filtros
func (r *userRepository) List(page, pageSize int, search string, status models.UserStatus, includeDeleted bool) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Validar parámetros de paginación
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	query := database.DB.Model(&models.User{})

	// Filtro por borrado lógico
	if !includeDeleted {
		query = query.Where("is_deleted = ?", false)
	}

	// Filtro por status
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Filtro de búsqueda
	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(email) LIKE ? OR LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(phone) LIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm,
		)
	}

	// Contar total de registros
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginación y ordenamiento
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdatePassword actualiza la contraseña de un usuario
func (r *userRepository) UpdatePassword(id uuid.UUID, newPasswordHash string) error {
	result := database.DB.Model(&models.User{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Update("password_hash", newPasswordHash)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}

	return nil
}

// VerifyEmail marca el email como verificado
func (r *userRepository) VerifyEmail(id uuid.UUID) error {
	result := database.DB.Model(&models.User{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Update("email_verified", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}

	return nil
}

// VerifyPhone marca el teléfono como verificado
func (r *userRepository) VerifyPhone(id uuid.UUID) error {
	result := database.DB.Model(&models.User{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Update("phone_verified", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}

	return nil
}

// EnableTwoFactor habilita la autenticación de dos factores
func (r *userRepository) EnableTwoFactor(id uuid.UUID) error {
	result := database.DB.Model(&models.User{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Update("two_factor_enabled", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}

	return nil
}

// DisableTwoFactor deshabilita la autenticación de dos factores
func (r *userRepository) DisableTwoFactor(id uuid.UUID) error {
	result := database.DB.Model(&models.User{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Update("two_factor_enabled", false)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}

	return nil
}

// UpdateLastLogin actualiza el timestamp del último login
func (r *userRepository) UpdateLastLogin(id uuid.UUID) error {
	now := time.Now()
	result := database.DB.Model(&models.User{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Update("last_login_at", now)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}

	return nil
}

// ExistsByEmail verifica si existe un usuario con el email dado
func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := database.DB.Model(&models.User{}).
		Where("email = ? AND is_deleted = ?", email, false).
		Count(&count).Error

	return count > 0, err
}

// GetActiveUsers obtiene todos los usuarios activos
func (r *userRepository) GetActiveUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Where("status = ? AND is_deleted = ?", models.UserStatusActive, false).
		Find(&users).Error

	return users, err
}

// RestoreUser restaura un usuario eliminado lógicamente
func (r *userRepository) RestoreUser(id uuid.UUID) error {
	var user models.User
	err := database.DB.Where("id = ? AND is_deleted = ?", id, true).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("deleted user not found")
		}
		return err
	}

	user.Restore()
	return database.DB.Save(&user).Error
}
