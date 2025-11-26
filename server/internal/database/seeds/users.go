package seeds

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"server/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// Estructuras para leer el YAML

type SeedUserRole struct {
	ApplicationClientID string `yaml:"application_client_id"`
	RoleName            string `yaml:"role_name"`
}

type SeedUser struct {
	ID     uuid.UUID      `yaml:"id"`
	Email  string         `yaml:"email"`
	DNI    string         `yaml:"dni"`
	Status string         `yaml:"status"`
	Roles  []SeedUserRole `yaml:"roles"`
}

// Seeder principal
func SeedUsersAndUserApplicationRoles(db *gorm.DB) error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding users y user_application_roles desde YAML...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	filePath := "internal/database/seeds/data/users.yml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo leer %s: %w", filePath, err)
	}

	var items []SeedUser
	if err := yaml.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("error al decodificar YAML de usuarios: %w", err)
	}

	for _, su := range items {
		email := strings.TrimSpace(su.Email)
		dni := strings.TrimSpace(su.DNI)

		if email == "" || dni == "" {
			logrus.Warnf("Usuario con email o DNI vacío en seed, se omite: email='%s', dni='%s'", email, dni)
			continue
		}

		status := strings.TrimSpace(su.Status)
		if status == "" {
			status = "active"
		}

		// 1) Buscar o crear usuario
		var user models.User
		err := db.Where("email = ?", email).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// -------------------------------
				// NUEVO: respetar ID si viene en el YAML
				// -------------------------------
				userID := su.ID
				if userID == uuid.Nil {
					userID = uuid.New()
				}

				now := time.Now()
				user = models.User{
					ID:        userID,
					Email:     email,
					DNI:       dni,
					Status:    status,
					CreatedAt: now,
					UpdatedAt: now,
					IsDeleted: false,
				}
				if err = db.Create(&user).Error; err != nil {
					return fmt.Errorf("error creando usuario '%s': %w", email, err)
				}

				if su.ID != uuid.Nil {
					logrus.Infof("Usuario creado con ID explícito: id='%s', email='%s', dni='%s'", userID, email, dni)
				} else {
					logrus.Infof("Usuario creado (ID generado): email='%s', dni='%s'", email, dni)
				}
			} else {
				return fmt.Errorf("error buscando usuario '%s': %w", email, err)
			}
		} else {
			// Usuario ya existe
			needsUpdate := false

			// (Opcional) podrías verificar si su.ID viene y es distinto de user.ID,
			// y solo loguear un warning:
			if su.ID != uuid.Nil && su.ID != user.ID {
				logrus.Warnf("SeedUser tiene ID='%s' pero en BD ya existe usuario '%s' con ID='%s'. Se mantiene el de BD.",
					su.ID, email, user.ID)
			}

			if user.DNI != dni {
				user.DNI = dni
				needsUpdate = true
			}
			if user.Status != status {
				user.Status = status
				needsUpdate = true
			}
			if needsUpdate {
				user.UpdatedAt = time.Now()
				if err := db.Save(&user).Error; err != nil {
					return fmt.Errorf("error actualizando usuario '%s': %w", email, err)
				}
				logrus.Infof("Usuario actualizado: email='%s', dni='%s'", email, dni)
			} else {
				logrus.Infof("Usuario ya existe (sin cambios): email='%s'", email)
			}
		}

		// 2) Procesar roles (igual que ya lo tenías)
		for _, sr := range su.Roles {
			appClientID := strings.TrimSpace(sr.ApplicationClientID)
			roleName := strings.TrimSpace(sr.RoleName)

			if appClientID == "" || roleName == "" {
				logrus.Warnf("Rol con datos incompletos para usuario '%s', se omite: app_client_id='%s', role='%s'",
					email, appClientID, roleName)
				continue
			}

			var app models.Application
			if err := db.Where("client_id = ?", appClientID).First(&app).Error; err != nil {
				return fmt.Errorf("no se encontró Application con client_id='%s' para usuario '%s': %w",
					appClientID, email, err)
			}

			var appRole models.ApplicationRole
			if err := db.Where("LOWER(name) = LOWER(?) AND application_id = ?", roleName, app.ID).
				First(&appRole).Error; err != nil {
				return fmt.Errorf("no se encontró ApplicationRole name='%s' en app client_id='%s' para usuario '%s': %w",
					roleName, appClientID, email, err)
			}

			var existing models.UserApplicationRole
			err := db.Where(
				"user_id = ? AND application_id = ? AND application_role_id = ?",
				user.ID, app.ID, appRole.ID,
			).First(&existing).Error

			if err == nil {
				logrus.Infof("UserApplicationRole ya existe: user='%s', app='%s', role='%s'",
					email, appClientID, roleName)
				continue
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("error verificando UserApplicationRole para user='%s', app='%s', role='%s': %w",
					email, appClientID, roleName, err)
			}

			now := time.Now()
			uRole := models.UserApplicationRole{
				ID:                uuid.New(),
				UserID:            user.ID,
				ApplicationID:     app.ID,
				ApplicationRoleID: appRole.ID,
				GrantedAt:         now,
				GrantedBy:         user.ID,
				IsDeleted:         false,
				CreatedAt:         now,
				UpdatedAt:         now,
			}

			if err := db.Create(&uRole).Error; err != nil {
				return fmt.Errorf("error creando UserApplicationRole para user='%s', app='%s', role='%s': %w",
					email, appClientID, roleName, err)
			}

			logrus.Infof("UserApplicationRole creado: user='%s', app='%s', role='%s'",
				email, appClientID, roleName)
		}
	}

	return nil
}
