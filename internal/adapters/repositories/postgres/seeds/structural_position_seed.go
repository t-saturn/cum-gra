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

type SeedPosition struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}

func SeedStructuralPositions() error {
	file, err := os.Open("internal/adapters/repositories/data/structural_positions.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer file.Close()

	var positions []SeedPosition
	if err := json.NewDecoder(file).Decode(&positions); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, p := range positions {
		var count int64
		err := database.DB.Model(&domain.StructuralPosition{}).
			Where("name = ? OR code = ?", p.Name, p.Code).
			Count(&count).Error
		if err != nil {
			return fmt.Errorf("error al verificar existencia de '%s': %w", p.Name, err)
		}

		if count > 0 {
			continue // ya existe, ignorar
		}

		position := domain.StructuralPosition{
			ID:          uuid.New(),
			Name:        p.Name,
			Code:        p.Code,
			Description: &p.Description,
			IsActive:    true,
			IsDeleted:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		position.Level = &p.Level

		if err := database.DB.Create(&position).Error; err != nil {
			return fmt.Errorf("error al insertar %s: %w", p.Name, err)
		}
	}

	return nil
}
