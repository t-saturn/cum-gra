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

type SeedOrganicUnit struct {
	Name        string `json:"name"`
	Acronym     string `json:"acronym"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
}

func SeedOrganicUnits() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding unidades orgánicas desde JSON...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	file, err := os.Open("internal/data/organic_units.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar el archivo: %v\n", cerr)
		}
	}()

	var units []SeedOrganicUnit
	if err := json.NewDecoder(file).Decode(&units); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, u := range units {
		var count int64
		err := config.DB.Model(&models.OrganicUnit{}).
			Where("name = ? OR acronym = ?", u.Name, u.Acronym).
			Count(&count).Error
		if err != nil {
			return fmt.Errorf("error al verificar existencia de '%s': %w", u.Name, err)
		}

		if count > 0 {
			logrus.Warnf("Unidad orgánica ya existe: %s (%s)", u.Name, u.Acronym)
			continue
		}

		unit := models.OrganicUnit{
			Name:        u.Name,
			Acronym:     u.Acronym,
			Brand:       &u.Brand,
			Description: &u.Description,
			IsActive:    true,
			IsDeleted:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := config.DB.Create(&unit).Error; err != nil {
			return fmt.Errorf("error al insertar unidad orgánica '%s': %w", u.Name, err)
		}

		logrus.Infof("Unidad orgánica insertada: %s (%s)", u.Name, u.Acronym)
	}

	return nil
}
