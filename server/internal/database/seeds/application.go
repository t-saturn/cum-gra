package seeds

import (
	"fmt"
	"os"
	"time"

	"server/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type SeedApplication struct {
	Name         string `yaml:"name"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Domain       string `yaml:"domain"`
	Logo         string `yaml:"logo"`
	Description  string `yaml:"description"`
	Status       string `yaml:"status"`

	// Por si más adelante quieres usarlos en el modelo:
	CallbackURLs []string `yaml:"callback_urls,omitempty"`
	IsFirstParty bool     `yaml:"is_first_party,omitempty"`
}

func SeedApplications(db *gorm.DB) error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding aplicaciones desde YAML...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	filePath := "internal/database/seeds/data/applications.yml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo YAML (%s): %w", filePath, err)
	}

	var apps []SeedApplication
	if err := yaml.Unmarshal(data, &apps); err != nil {
		return fmt.Errorf("error al decodificar YAML: %w", err)
	}

	for _, a := range apps {
		var count int64
		if err := db.Model(&models.Application{}).
			Where("client_id = ?", a.ClientID).
			Count(&count).Error; err != nil {
			return fmt.Errorf("error al verificar existencia de client_id '%s': %w", a.ClientID, err)
		}

		if count > 0 {
			logrus.Warnf("Aplicación ya existe: %s", a.Name)
			continue
		}

		logo := a.Logo
		desc := a.Description

		app := models.Application{
			ID:           uuid.New(),
			Name:         a.Name,
			ClientID:     a.ClientID,
			ClientSecret: a.ClientSecret,
			Domain:       a.Domain,
			Logo:         &logo,
			Description:  &desc,
			Status:       a.Status,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			IsDeleted:    false,
		}

		if err := db.Create(&app).Error; err != nil {
			return fmt.Errorf("error al insertar aplicación '%s': %w", a.Name, err)
		}

		logrus.Infof("Aplicación insertada: %s", a.Name)
	}

	return nil
}
