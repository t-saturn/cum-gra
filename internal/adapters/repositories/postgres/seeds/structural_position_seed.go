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

type SeedStructuralPosition struct {
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

	var entries []SeedStructuralPosition
	if err := json.NewDecoder(file).Decode(&entries); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, e := range entries {
		pos := domain.StructuralPosition{
			ID:          uuid.New(),
			Name:        e.Name,
			Code:        e.Code,
			Level:       intPtr(e.Level),
			Description: strPtr(e.Description),
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			IsDeleted:   false,
		}
		if err := database.DB.Create(&pos).Error; err != nil {
			return fmt.Errorf("error al insertar %s: %w", e.Name, err)
		}
	}

	return nil
}

func intPtr(i int) *int       { return &i }
func strPtr(s string) *string { return &s }
