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

type SeedPosition struct {
	Name        string `yaml:"name"`
	Code        string `yaml:"code"`
	Level       int    `yaml:"level"`
	Description string `yaml:"description"`
}

func SeedStructuralPositions(db *gorm.DB) error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding posiciones estructurales desde YAML...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	// Ruta sugerida para el YAML
	filePath := "internal/database/seeds/data/structural_positions.yml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo YAML (%s): %w", filePath, err)
	}

	var positions []SeedPosition
	if err := yaml.Unmarshal(data, &positions); err != nil {
		return fmt.Errorf("error al decodificar YAML: %w", err)
	}

	for _, p := range positions {
		var count int64
		err := db.Model(&models.StructuralPosition{}).
			Where("name = ? OR code = ?", p.Name, p.Code).
			Count(&count).Error
		if err != nil {
			return fmt.Errorf("error al verificar existencia de '%s': %w", p.Name, err)
		}

		if count > 0 {
			logrus.Warnf("Posición estructural ya existe: %s (%s)", p.Name, p.Code)
			continue
		}

		position := models.StructuralPosition{
			Name:        p.Name,
			Code:        p.Code,
			Description: &p.Description,
			IsActive:    true,
			IsDeleted:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Level es *int en el modelo
		position.Level = &p.Level

		if err := db.Create(&position).Error; err != nil {
			return fmt.Errorf("error al insertar posición estructural '%s': %w", p.Name, err)
		}

		logrus.Infof("Posición estructural insertada: %s (%s)", p.Name, p.Code)
	}

	return nil
}
