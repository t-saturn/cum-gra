package repositories

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	DNI                    string     `json:"dni"`
	Phone                  *string    `json:"phone,omitempty"`
	Status                 string     `json:"status"`
	IsDeleted              bool       `json:"-"`
	StructuralPositionID   *uuid.UUID `json:"structural_position_id,omitempty"`
	StructuralPositionName *string    `json:"structural_position,omitempty"`
	OrganicUnitID          *uuid.UUID `json:"organic_unit_id,omitempty"`
	OrganicUnitName        *string    `json:"organic_unit,omitempty"`
	Role                   string     `json:"role" gorm:"-"`
	ModulePermissions      []string   `json:"module_permissions" gorm:"-"`
	ModuleRestriccions     []string   `json:"module_restriccions" gorm:"-"`
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

// FindDataUser:
// - Trae la vista de usuario (estructura/unidad)
// - Trae roles (por appID)
// - Trae permisos de módulos (por appID) dados por sus roles (permission_type='access')
// - Trae restricciones por usuario (por appID) activas/no vencidas (limit_permission + denied)
func (r *UserRepository) FindDataUser(ctx context.Context, userID uuid.UUID, clientID string) (*UserOrgView, error) {
	// Activa salida de SQL para esta ruta (útil mientras depuras)
	db := r.db.WithContext(ctx).Debug()

	logrus.WithFields(logrus.Fields{
		"where":     "repo.FindDataUser.begin",
		"user_id":   userID,
		"client_id": clientID,
	}).Debug("inicio repo")

	var row UserOrgView
	// 1) Datos base del usuario
	if err := db.
		Table("users AS u").
		Select(`
            u.id,
            u.email,
            COALESCE(u.first_name, '') AS first_name,
            COALESCE(u.last_name, '')  AS last_name,
            u.dni,
            u.phone,
            u.status,
            u.is_deleted,
            u.structural_position_id,
            sp.name AS structural_position_name,
            u.organic_unit_id,
            ou.name AS organic_unit_name
        `).
		Joins("LEFT JOIN structural_positions sp ON sp.id = u.structural_position_id").
		Joins("LEFT JOIN organic_units ou ON ou.id = u.organic_unit_id").
		Where("u.id = ?", userID).
		Limit(1).
		Scan(&row).Error; err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":   "repo.FindDataUser.userBase",
			"user_id": userID,
		}).Error("falló query base de usuario")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"where":     "repo.FindDataUser.userBase",
		"email":     row.Email,
		"status":    row.Status,
		"isDeleted": row.IsDeleted,
	}).Debug("user base cargado")

	if row.IsDeleted {
		logrus.WithFields(logrus.Fields{
			"where": "repo.FindDataUser.userBase",
			"email": row.Email,
		}).Warn("usuario eliminado")
		return nil, ErrUserDeleted
	}
	if row.Status != "active" {
		logrus.WithFields(logrus.Fields{
			"where": "repo.FindDataUser.userBase",
			"email": row.Email,
			"st":    row.Status,
		}).Warn("usuario no activo")
		return nil, ErrUserDisabled
	}

	// 2) Rol único
	var roleName string
	if err := db.
		Table("user_application_roles AS uar").
		Select("LOWER(ar.name) AS name").
		Joins("JOIN application_roles ar ON ar.id = uar.application_role_id").
		Joins("JOIN applications a ON a.id = uar.application_id AND a.client_id = ?", clientID).
		Where("uar.user_id = ? AND uar.is_deleted = false AND uar.revoked_at IS NULL", userID).
		Limit(1).
		Row().
		Scan(&roleName); err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":     "repo.FindDataUser.role",
			"user_id":   userID,
			"client_id": clientID,
		}).Error("falló obtener rol")
		return nil, err
	}
	row.Role = roleName
	logrus.WithFields(logrus.Fields{
		"where": "repo.FindDataUser.role",
		"role":  row.Role,
	}).Debug("rol resuelto")

	// 3) Permisos
	var permNames []string
	if err := db.
		Table("users AS u").
		Select("DISTINCT m.name AS module_name").
		Joins("JOIN user_application_roles uar ON u.id = uar.user_id").
		Joins("JOIN application_roles ar ON uar.application_role_id = ar.id").
		Joins("JOIN applications a ON ar.application_id = a.id").
		Joins("JOIN module_role_permissions mrp ON ar.id = mrp.application_role_id AND mrp.is_deleted = false").
		Joins("JOIN modules m ON mrp.module_id = m.id").
		Where(`
        a.client_id = ? AND
        u.id = ? AND
        u.is_deleted = false AND
        a.is_deleted = false AND
        ar.is_deleted = false AND
        m.is_deleted = false AND
        m.status = 'active' AND
        uar.is_deleted = false AND
        uar.revoked_at IS NULL
    `, clientID, userID).
		Pluck("module_name", &permNames).Error; err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":     "repo.FindDataUser.perms",
			"user_id":   userID,
			"client_id": clientID,
			"role":      roleName, // si quieres loguearlo, aunque aquí no filtramos por nombre de rol
		}).Error("falló obtener permisos (por nombre)")
		return nil, err
	}
	row.ModulePermissions = permNames

	logrus.WithFields(logrus.Fields{
		"where":      "repo.FindDataUser.perms",
		"perm_count": len(row.ModulePermissions),
		"sample":     row.ModulePermissions,
	}).Debug("permisos (nombres) resueltos")

	// 4. Restricciones por client_id (limit_permission + denied + no vencidas) → devolver NOMBRES
	var restrNames []string
	now := time.Now()

	if err := db.
		Table("user_module_restrictions AS umr").
		Select("DISTINCT m.name AS module_name").
		Joins("JOIN applications a ON a.id = umr.application_id AND a.client_id = ?", clientID).
		Joins("JOIN modules m ON m.id = umr.module_id").
		Where(`
        umr.user_id = ? AND
        umr.is_deleted = false AND
        (umr.expires_at IS NULL OR umr.expires_at > ?) AND
        m.status = 'active' AND
        (a.is_deleted = false OR a.is_deleted IS NULL) AND
        (m.is_deleted = false OR m.is_deleted IS NULL) AND
        LOWER(umr.restriction_type) = 'limit_permission' AND
        LOWER(COALESCE(umr.max_permission_level, 'denied')) = 'denied'
    `, userID, now).
		Order("module_name ASC"). // seguro con DISTINCT porque ordena por el mismo select
		Pluck("module_name", &restrNames).Error; err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":     "repo.FindDataUser.restr",
			"user_id":   userID,
			"client_id": clientID,
		}).Error("falló obtener restricciones (por nombre)")
		return nil, err
	}

	row.ModuleRestriccions = restrNames
	logrus.WithFields(logrus.Fields{
		"where":       "repo.FindDataUser.restr",
		"restr_count": len(row.ModuleRestriccions),
		"sample":      row.ModuleRestriccions,
	}).Debug("restricciones (nombres) resueltas")

	logrus.WithFields(logrus.Fields{
		"where":       "repo.FindDataUser.end",
		"email":       row.Email,
		"role":        row.Role,
		"perm_count":  len(row.ModulePermissions),
		"restr_count": len(row.ModuleRestriccions),
	}).Info("FindDataUser OK")

	return &row, nil
}
