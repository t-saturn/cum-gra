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

type SeedModule struct {
	Item            *string `json:"item"`
	Name            string  `json:"name"`
	Route           *string `json:"route"`
	Icon            *string `json:"icon"`
	ParentName      *string `json:"parent_name"`
	ApplicationName string  `json:"application_name"`
	SortOrder       int     `json:"sort_order"`
	Status          string  `json:"status"`
}

func SeedModules() error {
	file, err := os.Open("internal/adapters/repositories/data/modules.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer file.Close()

	var modules []SeedModule
	if err := json.NewDecoder(file).Decode(&modules); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, m := range modules {
		// Buscar la aplicación por nombre
		var app domain.Application
		err := database.DB.Where("name = ?", m.ApplicationName).First(&app).Error
		if err != nil {
			return fmt.Errorf("no se encontró la aplicación '%s': %w", m.ApplicationName, err)
		}

		// Obtener ID del módulo padre si existe
		var parentID *uuid.UUID
		if m.ParentName != nil {
			var parent domain.Module
			err := database.DB.Where("name = ?", *m.ParentName).First(&parent).Error
			if err != nil {
				return fmt.Errorf("no se encontró el módulo padre '%s': %w", *m.ParentName, err)
			}
			parentID = &parent.ID
		}

		// Evitar duplicados por nombre
		var existing domain.Module
		if err := database.DB.Where("name = ?", m.Name).First(&existing).Error; err == nil {
			continue // ya existe
		}

		module := domain.Module{
			ID:            uuid.New(),
			Item:          m.Item,
			Name:          m.Name,
			Route:         m.Route,
			Icon:          m.Icon,
			ParentID:      parentID,
			ApplicationID: &app.ID,
			SortOrder:     m.SortOrder,
			Status:        m.Status,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := database.DB.Create(&module).Error; err != nil {
			return fmt.Errorf("error al insertar módulo '%s': %w", m.Name, err)
		}
	}

	return nil
}
