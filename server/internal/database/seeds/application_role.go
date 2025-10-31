package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"server/internal/config"
	"server/internal/models"
)

type SeedApplicationRole struct {
	Name                string  `json:"name"`
	Description         *string `json:"description"`
	ApplicationClientID string  `json:"application_client_id"` // preferido para resolver la app
	ApplicationName     string  `json:"application_name"`      // fallback si no hay client_id
}

func SeedApplicationRoles() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding application_roles desde JSON...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	f, err := os.Open("internal/data/application_roles.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir internal/data/application_roles.json: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar archivo: %v\n", cerr)
		}
	}()

	var input []SeedApplicationRole
	if err := json.NewDecoder(f).Decode(&input); err != nil {
		return fmt.Errorf("error al decodificar JSON de roles: %w", err)
	}

	for _, r := range input {
		var app models.Application
		var qErr error

		if strings.TrimSpace(r.ApplicationClientID) != "" {
			qErr = config.DB.
				Where("client_id = ?", r.ApplicationClientID).
				First(&app).Error
		} else {
			qErr = config.DB.
				Where("LOWER(name) = LOWER(?)", r.ApplicationName).
				First(&app).Error
		}
		if qErr != nil {
			return fmt.Errorf("no se encontró la aplicación (client_id='%s', name='%s'): %w",
				r.ApplicationClientID, r.ApplicationName, qErr)
		}

		var count int64
		if err := config.DB.Model(&models.ApplicationRole{}).
			Where("LOWER(name) = LOWER(?) AND application_id = ?", r.Name, app.ID).
			Count(&count).Error; err != nil {
			return fmt.Errorf("error verificando duplicados para role '%s' en app '%s': %w",
				r.Name, app.Name, err)
		}
		var desc *string
		if r.Description != nil && strings.TrimSpace(*r.Description) != "" {
			desc = r.Description
		}

		if count > 0 && desc != nil {
			if err := config.DB.Model(&models.ApplicationRole{}).
				Where("LOWER(name) = LOWER(?) AND application_id = ?", r.Name, app.ID).
				Updates(map[string]any{"description": desc, "updated_at": time.Now()}).Error; err != nil {
				return fmt.Errorf("error actualizando descripción de role '%s' app '%s': %w", r.Name, app.Name, err)
			}
			logrus.Infof("ApplicationRole actualizado: role='%s' app='%s'", r.Name, app.Name)
			continue
		}

		role := models.ApplicationRole{
			ID:            uuid.New(),
			Name:          r.Name,
			Description:   desc,
			ApplicationID: app.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			IsDeleted:     false,
		}

		if err := config.DB.Create(&role).Error; err != nil {
			return fmt.Errorf("error al insertar role '%s' en app '%s': %w", r.Name, app.Name, err)
		}

		logrus.Infof("ApplicationRole insertado: role='%s' app='%s'", r.Name, app.Name)
	}

	return nil
}
