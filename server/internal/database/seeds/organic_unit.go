package seeds

import (
	"fmt"
	"os"
	"time"

	"server/internal/models"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type SeedOrganicUnit struct {
	Name        string `yaml:"name"`
	Acronym     string `yaml:"acronym"`
	Brand       string `yaml:"brand"`
	Description string `yaml:"description"`
}

func SeedOrganicUnits(db *gorm.DB) error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding unidades orgánicas desde YAML...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	filePath := "internal/database/seeds/data/organic_units.yml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo YAML (%s): %w", filePath, err)
	}

	var units []SeedOrganicUnit
	if err := yaml.Unmarshal(data, &units); err != nil {
		return fmt.Errorf("error al decodificar YAML: %w", err)
	}

	for _, u := range units {
		var count int64
		err := db.Model(&models.OrganicUnit{}).
			Where("name = ? OR acronym = ?", u.Name, u.Acronym).
			Count(&count).Error
		if err != nil {
			return fmt.Errorf("error al verificar existencia de '%s': %w", u.Name, err)
		}

		if count > 0 {
			logrus.Warnf("Unidad orgánica ya existe: %s (%s)", u.Name, u.Acronym)
			continue
		}

		// Punteros para Brand y Description (porque en el modelo son *string)
		var brand *string
		var description *string

		if u.Brand != "" {
			b := u.Brand
			brand = &b
		}
		if u.Description != "" {
			d := u.Description
			description = &d
		}

		unit := models.OrganicUnit{
			Name:        u.Name,
			Acronym:     u.Acronym,
			Brand:       brand,
			Description: description,
			IsActive:    true,
			IsDeleted:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := db.Create(&unit).Error; err != nil {
			return fmt.Errorf("error al insertar unidad orgánica '%s': %w", u.Name, err)
		}

		logrus.Infof("Unidad orgánica insertada: %s (%s)", u.Name, u.Acronym)
	}

	return nil
}
