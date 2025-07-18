package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/t-saturn/central-user-manager/config"
	"github.com/t-saturn/central-user-manager/internal/models"
)

// SeedModule representa la estructura de un módulo utilizada para poblar la base de datos desde un archivo JSON.
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

// SeedModules inserta registros de módulos en la base de datos desde un archivo JSON, asociándolos a sus aplicaciones y padres si corresponde.
func SeedModules() error {
	logrus.Info("Seeding módulos desde JSON...")

	file, err := os.Open("data/modules.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar el archivo: %v\n", cerr)
		}
	}()

	var seedData []SeedModule
	if err := json.NewDecoder(file).Decode(&seedData); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	db := config.DB

	// Indexar módulos insertados por nombre para posterior asignación de ParentID
	inserted := make(map[string]uuid.UUID)

	for _, sm := range seedData {
		var app models.Application
		err := db.Where("LOWER(name) = ?", strings.ToLower(sm.ApplicationName)).First(&app).Error
		if err != nil {
			return fmt.Errorf("no se encontró la aplicación '%s': %w", sm.ApplicationName, err)
		}

		var parentID *uuid.UUID // nil por defecto, no es necesario asignarlo explícitamente
		if sm.ParentName != nil {
			if pid, ok := inserted[*sm.ParentName]; ok {
				parentID = &pid
			} else {
				var parent models.Module
				err := db.Where("name = ? AND application_id = ?", *sm.ParentName, app.ID).First(&parent).Error
				if err != nil {
					return fmt.Errorf("no se encontró el módulo padre '%s': %w", *sm.ParentName, err)
				}
				parentID = &parent.ID
				inserted[parent.Name] = parent.ID
			}
		}

		// Verificar existencia correctamente
		var count int64
		query := db.Model(&models.Module{}).Where("name = ? AND application_id = ?", sm.Name, app.ID)
		if parentID == nil {
			query = query.Where("parent_id IS NULL")
		} else {
			query = query.Where("parent_id = ?", *parentID)
		}
		if err := query.Count(&count).Error; err != nil {
			return fmt.Errorf("error verificando existencia del módulo '%s': %w", sm.Name, err)
		}
		if count > 0 {
			logrus.Warnf("Módulo ya existe: %s", sm.Name)
			continue
		}

		module := models.Module{
			ID:            uuid.New(),
			Item:          sm.Item,
			Name:          sm.Name,
			Route:         sm.Route,
			Icon:          sm.Icon,
			ParentID:      parentID,
			ApplicationID: &app.ID,
			SortOrder:     sm.SortOrder,
			Status:        sm.Status,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := db.Create(&module).Error; err != nil {
			return fmt.Errorf("error al insertar módulo '%s': %w", sm.Name, err)
		}

		inserted[module.Name] = module.ID
		logrus.Infof("Módulo insertado: %s", sm.Name)
	}

	return nil
}
