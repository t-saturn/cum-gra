package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/t-saturn/central-user-manager/internal/config"
	"github.com/t-saturn/central-user-manager/internal/models"
	"github.com/t-saturn/central-user-manager/pkg/security"
)

// -----------------------------------------------------------------------------
// JSON structs esperados
// -----------------------------------------------------------------------------

type SeedUser struct {
	ID                   *string `json:"id,omitempty"`
	Email                string  `json:"email"`
	Password             string  `json:"password"`
	FirstName            *string `json:"first_name"`
	LastName             *string `json:"last_name"`
	Phone                *string `json:"phone"`
	DNI                  string  `json:"dni"`
	Status               string  `json:"status"`
	StructuralPositionID *string `json:"structural_position_id"`
	OrganicUnitID        *string `json:"organic_unit_id"`
	// Cualquier extra (p.ej. user_application_role_ids) será ignorado
}

type SeedUserApplicationRole struct {
	ID                  *string `json:"id,omitempty"`
	UserEmail           string  `json:"user_email"`
	ApplicationClientID string  `json:"application_client_id,omitempty"` // preferido
	ApplicationName     string  `json:"application_name,omitempty"`      // fallback
	ApplicationRoleName string  `json:"application_role_name"`
	GrantedByEmail      string  `json:"granted_by_email"`
}

// -----------------------------------------------------------------------------
// Seeder maestro (usuarios primero, luego relaciones de rol)
// -----------------------------------------------------------------------------

func SeedUsersAndUserApplicationRoles() error {
	logrus.Info("==============================================================================================")
	logrus.Info("Seeding: USERS y luego USER_APPLICATION_ROLES (agregando relación al usuario)")
	logrus.Info("==============================================================================================")

	if err := seedUsersOnly(); err != nil {
		return err
	}
	if err := linkUserApplicationRoles(); err != nil {
		return err
	}

	logrus.Info("Seeding completo: USERS y USER_APPLICATION_ROLES")
	return nil
}

// -----------------------------------------------------------------------------
// Paso 1: Insertar Users (sin relaciones)
// -----------------------------------------------------------------------------

func seedUsersOnly() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding users desde JSON...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	f, err := os.Open("data/users.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir data/users.json: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar users.json: %v\n", cerr)
		}
	}()

	var input []SeedUser
	if err := json.NewDecoder(f).Decode(&input); err != nil {
		return fmt.Errorf("error al decodificar JSON de users: %w", err)
	}

	for _, u := range input {
		// Evitar duplicados por email o dni
		var count int64
		if err := config.DB.Model(&models.User{}).
			Where("LOWER(email) = LOWER(?) OR dni = ?", u.Email, u.DNI).
			Count(&count).Error; err != nil {
			return fmt.Errorf("error verificando duplicados de user '%s': %w", u.Email, err)
		}
		if count > 0 {
			logrus.Warnf("User ya existe: email='%s' dni='%s'", u.Email, u.DNI)
			continue
		}

		// Parse IDs opcionales
		var id uuid.UUID
		if u.ID != nil && strings.TrimSpace(*u.ID) != "" {
			parsed, pErr := uuid.Parse(*u.ID)
			if pErr != nil {
				return fmt.Errorf("user '%s' id inválido '%s': %w", u.Email, *u.ID, pErr)
			}
			id = parsed
		} else {
			id = uuid.New()
		}

		var spID *uuid.UUID
		if u.StructuralPositionID != nil && strings.TrimSpace(*u.StructuralPositionID) != "" {
			p, perr := uuid.Parse(*u.StructuralPositionID)
			if perr != nil {
				return fmt.Errorf("user '%s' structural_position_id inválido '%s': %w", u.Email, *u.StructuralPositionID, perr)
			}
			spID = &p
		}

		var ouID *uuid.UUID
		if u.OrganicUnitID != nil && strings.TrimSpace(*u.OrganicUnitID) != "" {
			p, perr := uuid.Parse(*u.OrganicUnitID)
			if perr != nil {
				return fmt.Errorf("user '%s' organic_unit_id inválido '%s': %w", u.Email, *u.OrganicUnitID, perr)
			}
			ouID = &p
		}

		argon := security.NewArgon2Service()

		hash, err := argon.HashPassword(u.Password)
		if err != nil {
			return fmt.Errorf("error generando hash para user '%s': %w", u.Email, err)
		}

		user := models.User{
			ID:                   id,
			Email:                u.Email,
			PasswordHash:         hash,
			FirstName:            u.FirstName,
			LastName:             u.LastName,
			Phone:                u.Phone,
			DNI:                  u.DNI,
			Status:               u.Status,
			StructuralPositionID: spID,
			OrganicUnitID:        ouID,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
			IsDeleted:            false,
		}

		if err := config.DB.Create(&user).Error; err != nil {
			return fmt.Errorf("error insertando user '%s': %w", u.Email, err)
		}

		logrus.Infof("User insertado: email='%s'", u.Email)
	}

	return nil
}

// -----------------------------------------------------------------------------
// Paso 2: Insertar UserApplicationRoles y "agregar la relación" al usuario
// -----------------------------------------------------------------------------

func linkUserApplicationRoles() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding user_application_roles desde JSON y agregando asociación al usuario...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	f, err := os.Open("data/user_application_roles.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir data/user_application_roles.json: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar user_application_roles.json: %v\n", cerr)
		}
	}()

	var input []SeedUserApplicationRole
	if err := json.NewDecoder(f).Decode(&input); err != nil {
		return fmt.Errorf("error al decodificar JSON de user_application_roles: %w", err)
	}

	for _, r := range input {
		// 1) User
		var user models.User
		if err := config.DB.Where("LOWER(email) = LOWER(?)", r.UserEmail).First(&user).Error; err != nil {
			return fmt.Errorf("no se encontró user con email '%s': %w", r.UserEmail, err)
		}

		// 2) GrantedBy (si no existe, usar el mismo user)
		var grantedBy models.User
		if err := config.DB.Where("LOWER(email) = LOWER(?)", r.GrantedByEmail).First(&grantedBy).Error; err != nil {
			grantedBy = user
			logrus.Warnf("No se encontró granted_by '%s', usando el mismo user '%s'",
				r.GrantedByEmail, r.UserEmail)
		}

		// 3) Application
		var app models.Application
		if strings.TrimSpace(r.ApplicationClientID) != "" {
			if err := config.DB.Where("client_id = ?", r.ApplicationClientID).First(&app).Error; err != nil {
				return fmt.Errorf("no se encontró application client_id='%s': %w", r.ApplicationClientID, err)
			}
		} else {
			if err := config.DB.Where("LOWER(name) = LOWER(?)", r.ApplicationName).First(&app).Error; err != nil {
				return fmt.Errorf("no se encontró application name='%s': %w", r.ApplicationName, err)
			}
		}

		// 4) Role en esa app
		var role models.ApplicationRole
		if err := config.DB.Where("LOWER(name) = LOWER(?) AND application_id = ?", r.ApplicationRoleName, app.ID).First(&role).Error; err != nil {
			return fmt.Errorf("no se encontró role '%s' en app '%s': %w", r.ApplicationRoleName, app.Name, err)
		}

		// 5) Evitar duplicado activo exacto
		var count int64
		if err := config.DB.Model(&models.UserApplicationRole{}).
			Where("user_id = ? AND application_id = ? AND application_role_id = ? AND is_deleted = false AND revoked_at IS NULL",
				user.ID, app.ID, role.ID).
			Count(&count).Error; err != nil {
			return fmt.Errorf("error verificando duplicado de UAR (user='%s', app='%s', role='%s'): %w",
				user.Email, app.Name, role.Name, err)
		}
		if count > 0 {
			logrus.Warnf("UserApplicationRole ya existe: user='%s' app='%s' role='%s'",
				user.Email, app.Name, role.Name)
			continue
		}

		// 6) ID opcional
		var id uuid.UUID
		if r.ID != nil && strings.TrimSpace(*r.ID) != "" {
			parsed, pErr := uuid.Parse(*r.ID)
			if pErr != nil {
				return fmt.Errorf("UAR id inválido '%s' (user='%s', app='%s', role='%s'): %w",
					*r.ID, user.Email, app.Name, role.Name, pErr)
			}
			id = parsed
		} else {
			id = uuid.New()
		}

		now := time.Now()

		// 7) Crear UAR y "agregar relación" al usuario vía Association (setea user_id)
		uar := models.UserApplicationRole{
			ID:                id,
			ApplicationID:     app.ID,
			ApplicationRoleID: role.ID,
			GrantedAt:         now,
			GrantedBy:         grantedBy.ID,
			IsDeleted:         false,
		}

		// Usamos la asociación para que GORM asigne user_id y cree el registro
		if err := config.DB.Model(&user).Association("UserApplicationRoles").Append(&uar); err != nil {
			return fmt.Errorf("error asociando UAR con user '%s': %w", user.Email, err)
		}

		// 8) (Opcional) Tocar updated_at del usuario tras asignar rol
		if err := config.DB.Model(&models.User{}).
			Where("id = ?", user.ID).
			Update("updated_at", now).Error; err != nil {
			return fmt.Errorf("error actualizando updated_at de user '%s': %w", user.Email, err)
		}

		logrus.Infof("UAR insertado y asociado: user='%s' app='%s' role='%s'", user.Email, app.Name, role.Name)
	}

	return nil
}
