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

// ─────────────────────────────────────────────────────────────────────────────
// MODELOS / PROYECCIONES
// ─────────────────────────────────────────────────────────────────────────────

type User struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primaryKey;column:id"`
	Email                string     `gorm:"column:email"`
	PasswordHash         string     `gorm:"column:password_hash"`
	DNI                  string     `gorm:"column:dni"`
	Status               string     `gorm:"column:status"`
	IsDeleted            bool       `gorm:"column:is_deleted"`
	StructuralPositionID *uuid.UUID `gorm:"type:uuid;column:structural_position_id"`
	OrganicUnitID        *uuid.UUID `gorm:"type:uuid;column:organic_unit_id"`
}

func (User) TableName() string { return "users" }

// Proyección para devolver nombres de posición/unidad + nombre completo
type UserOrgView struct {
	ID                     uuid.UUID  `json:"id"`
	Email                  string     `json:"email"`
	FirstName              string     `json:"first_name"`
	LastName               string     `json:"last_name"`
	DNI                    string     `json:"dni"`             // <- NUEVO
	Phone                  *string    `json:"phone,omitempty"` // <- NUEVO (nullable)
	Status                 string     `json:"status"`
	IsDeleted              bool       `json:"-"`
	StructuralPositionID   *uuid.UUID `json:"structural_position_id,omitempty"`
	StructuralPositionName *string    `json:"structural_position,omitempty"`
	OrganicUnitID          *uuid.UUID `json:"organic_unit_id,omitempty"`
	OrganicUnitName        *string    `json:"organic_unit,omitempty"`
}

// (Opcional) Si usas los modelos GORM de estas tablas en otro lado:
type StructuralPosition struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;column:id"`
	Name string    `gorm:"column:name"`
}

func (StructuralPosition) TableName() string { return "structural_positions" }

type OrganicUnit struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;column:id"`
	Name string    `gorm:"column:name"`
}

func (OrganicUnit) TableName() string { return "organic_units" }

// ─────────────────────────────────────────────────────────────────────────────
// QUERIES
// ─────────────────────────────────────────────────────────────────────────────

// FindByID simple (opcional, por si te sirve en otros flujos)
func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	var u User
	err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

// FindActiveByEmailOrDNI busca un usuario por email o DNI (tu implementación existente, sin cambios)
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

// FindActiveByIDWithOrgNames busca por ID y devuelve nombres de posición estructural y unidad orgánica.
// - Valida usuario eliminado o deshabilitado.
// - Hace LEFT JOIN con structural_positions y organic_units para traer los nombres.
func (r *UserRepository) FindActiveByIDWithOrgNames(ctx context.Context, id uuid.UUID) (*UserOrgView, error) {
	var row UserOrgView

	q := r.db.WithContext(ctx).
		Table("users AS u").
		Select(`
            u.id,
            u.email,
            u.first_name,
            u.last_name,
            u.dni,            -- <- NUEVO
            u.phone,          -- <- NUEVO
            u.status,
            u.is_deleted,
            u.structural_position_id,
            sp.name AS structural_position_name,
            u.organic_unit_id,
            ou.name AS organic_unit_name
        `).
		Joins("LEFT JOIN structural_positions sp ON sp.id = u.structural_position_id").
		Joins("LEFT JOIN organic_units ou ON ou.id = u.organic_unit_id").
		Where("u.id = ?", id).
		Limit(1)

	if err := q.Scan(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if row.IsDeleted {
		return nil, ErrUserDeleted
	}
	if row.Status != "active" {
		return nil, ErrUserDisabled
	}

	return &row, nil
}
