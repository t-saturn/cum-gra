package seeds

import (
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

type SeedApplicationRole struct {
	Name                string  `yaml:"name"`
	Description         *string `yaml:"description"`
	ApplicationClientID string  `yaml:"application_client_id"`
}

func SeedApplicationRoles(db *gorm.DB) error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding application_roles desde YAML...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	filePath := "internal/database/seeds/data/application_roles.yml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo leer %s: %w", filePath, err)
	}

	var input []SeedApplicationRole
	if err := yaml.Unmarshal(data, &input); err != nil {
		return fmt.Errorf("error al decodificar YAML de roles: %w", err)
	}

	for _, r := range input {
		if strings.TrimSpace(r.ApplicationClientID) == "" {
			logrus.Warnf("SeedApplicationRole sin application_client_id, se omite: role='%s'", r.Name)
			continue
		}

		// Buscar la aplicación por client_id
		var app models.Application
		if err := db.
			Where("client_id = ?", r.ApplicationClientID).
			First(&app).Error; err != nil {
			return fmt.Errorf("no se encontró Application con client_id='%s' para role='%s': %w",
				r.ApplicationClientID, r.Name, err)
		}

		// Verificar si el rol ya existe para esa aplicación
		var count int64
		if err := db.Model(&models.ApplicationRole{}).
			Where("LOWER(name) = LOWER(?) AND application_id = ?", r.Name, app.ID).
			Count(&count).Error; err != nil {
			return fmt.Errorf("error verificando duplicados para role '%s' en app '%s': %w",
				r.Name, app.Name, err)
		}

		// Normalizar descripción
		var desc *string
		if r.Description != nil && strings.TrimSpace(*r.Description) != "" {
			d := strings.TrimSpace(*r.Description)
			desc = &d
		}

		// Si ya existe y hay descripción nueva, actualizamos
		if count > 0 && desc != nil {
			if err := db.Model(&models.ApplicationRole{}).
				Where("LOWER(name) = LOWER(?) AND application_id = ?", r.Name, app.ID).
				Updates(map[string]any{
					"description": desc,
					"updated_at":  time.Now(),
				}).Error; err != nil {
				return fmt.Errorf("error actualizando descripción de role '%s' app '%s': %w",
					r.Name, app.Name, err)
			}
			logrus.Infof("ApplicationRole actualizado: role='%s' app='%s'", r.Name, app.Name)
			continue
		}

		// Crear nuevo rol
		role := models.ApplicationRole{
			ID:            uuid.New(),
			Name:          r.Name,
			Description:   desc,
			ApplicationID: app.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			IsDeleted:     false,
		}

		if err := db.Create(&role).Error; err != nil {
			return fmt.Errorf("error al insertar role '%s' en app '%s': %w", r.Name, app.Name, err)
		}

		logrus.Infof("ApplicationRole insertado: role='%s' app='%s'", r.Name, app.Name)
	}

	return nil
}
