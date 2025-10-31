package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/sirupsen/logrus"
)

type SeedPosition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}

func SeedStructuralPositions() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding posiciones estructurales desde JSON...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	file, err := os.Open("internal/data/structural_positions.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar el archivo: %v\n", cerr)
		}
	}()

	var positions []SeedPosition
	if err := json.NewDecoder(file).Decode(&positions); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, p := range positions {
		var count int64
		err := config.DB.Model(&models.StructuralPosition{}).
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

		position.Level = &p.Level

		if err := config.DB.Create(&position).Error; err != nil {
			return fmt.Errorf("error al insertar posición estructural '%s': %w", p.Name, err)
		}

		logrus.Infof("Posición estructural insertada: %s (%s)", p.Name, p.Code)
	}

	return nil
}
