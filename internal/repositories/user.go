package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("usuario no encontrado")
	ErrUserDeleted  = errors.New("este usuario está eliminado")
	ErrUserDisabled = errors.New("el usuario está deshabilitado")
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository construye un UserRepository usando la conexión de GORM.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	DNI          string
	Status       string
	IsDeleted    bool
}

// FindActiveByEmailOrDNI busca un usuario por email o DNI, sin filtrar is_deleted
// para poder detectar usuarios eliminados. Después de cargarlo, comprueba:
//   - si user.IsDeleted == true   → ErrUserDeleted
//   - si user.Status != "active"  → ErrUserDisabled
//   - si no existe                → ErrUserNotFound
func (r *UserRepository) FindActiveByEmailOrDNI(ctx context.Context, email, dni *string) (*User, error) {
	var user User

	q := r.db.WithContext(ctx).Model(&User{})

	// Construir la condición OR
	switch {
	case email != nil && dni != nil:
		q = q.Where("email = ? OR dni = ?", *email, *dni)
	case email != nil:
		q = q.Where("email = ?", *email)
	case dni != nil:
		q = q.Where("dni = ?", *dni)
	default:
		return nil, errors.New("se requiere email o dni")
	}

	// No filtramos is_deleted aquí para poder manejarlo manualmente
	err := q.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Validaciones post-consulta
	if user.IsDeleted {
		return nil, ErrUserDeleted
	}
	if user.Status != "active" {
		return nil, ErrUserDisabled
	}

	return &user, nil
}
