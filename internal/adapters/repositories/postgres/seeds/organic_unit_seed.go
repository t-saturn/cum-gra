package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
)

type SeedOrganicUnit struct {
	Name        string `json:"name"`
	Acronym     string `json:"acronym"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
}

func SeedOrganicUnits() error {
	file, err := os.Open("internal/adapters/repositories/data/organic_units.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer file.Close()

	var units []SeedOrganicUnit
	if err := json.NewDecoder(file).Decode(&units); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, u := range units {
		var count int64
		err := database.DB.Model(&domain.OrganicUnit{}).
			Where("name = ? OR acronym = ?", u.Name, u.Acronym).
			Count(&count).Error
		if err != nil {
			return fmt.Errorf("error al verificar existencia de '%s': %w", u.Name, err)
		}

		if count > 0 {
			continue
		}

		unit := domain.OrganicUnit{
			ID:          uuid.New(),
			Name:        u.Name,
			Acronym:     u.Acronym,
			Brand:       &u.Brand,
			Description: &u.Description,
			IsActive:    true,
			IsDeleted:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := database.DB.Create(&unit).Error; err != nil {
			return fmt.Errorf("error al insertar unidad orgánica '%s': %w", u.Name, err)
		}
	}

	return nil
}
