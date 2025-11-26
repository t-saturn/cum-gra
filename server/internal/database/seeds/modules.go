package seeds

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"server/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// Estructura para leer modules.yml
type SeedModule struct {
	ApplicationClientID string   `yaml:"application_client_id"`
	Item                string   `yaml:"item"`         // Grupo: Menú, Gestión, Seguridad, Configuración
	Name                string   `yaml:"name"`         // Nombre del módulo
	Route               string   `yaml:"route"`        // Ruta del módulo
	Icon                string   `yaml:"icon"`         // Nombre del ícono (solo referencial)
	ParentRoute         string   `yaml:"parent_route"` // Ruta del módulo padre (opcional)
	SortOrder           int      `yaml:"sort_order"`   // Orden entre hermanos
	Status              string   `yaml:"status"`       // active / inactive
	Roles               []string `yaml:"roles"`        // Nombres de roles que tienen acceso
}

// Seeder principal
func SeedModules(db *gorm.DB) error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding modules y module_role_permissions desde YAML...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	filePath := "internal/database/seeds/data/modules.yml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo leer %s: %w", filePath, err)
	}

	var items []SeedModule
	if err := yaml.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("error al decodificar YAML de módulos: %w", err)
	}

	for _, sm := range items {
		appClientID := strings.TrimSpace(sm.ApplicationClientID)
		if appClientID == "" {
			logrus.Warnf("SeedModule sin application_client_id, se omite: name='%s', route='%s'", sm.Name, sm.Route)
			continue
		}

		// 1) Buscar aplicación por client_id
		var app models.Application
		if err := db.Where("client_id = ?", appClientID).First(&app).Error; err != nil {
			return fmt.Errorf("no se encontró Application con client_id='%s' para módulo '%s': %w",
				appClientID, sm.Name, err)
		}

		// 2) Resolver ParentID si hay parent_route
		var parentID *uuid.UUID
		parentRoute := strings.TrimSpace(sm.ParentRoute)
		if parentRoute != "" {
			var parentModule models.Module
			if err := db.
				Where("application_id = ? AND route = ?", app.ID, parentRoute).
				First(&parentModule).Error; err != nil {
				return fmt.Errorf("no se encontró módulo padre route='%s' para módulo '%s' en app '%s': %w",
					parentRoute, sm.Name, appClientID, err)
			}
			pid := parentModule.ID
			parentID = &pid
		}

		// 3) Buscar si el módulo ya existe por (application_id, route)
		var module models.Module
		err := db.
			Where("application_id = ? AND route = ?", app.ID, sm.Route).
			First(&module).Error

		now := time.Now()
		item := strings.TrimSpace(sm.Item)
		if item == "" {
			item = "Menú"
		}
		status := strings.TrimSpace(sm.Status)
		if status == "" {
			status = "active"
		}

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("error buscando módulo route='%s' app='%s': %w",
					sm.Route, appClientID, err)
			}

			// 3.1) Crear módulo nuevo
			itemCopy := item
			iconCopy := strings.TrimSpace(sm.Icon)
			appIDCopy := app.ID

			module = models.Module{
				ID:            uuid.New(),
				Item:          &itemCopy,
				Name:          sm.Name,
				Route:         strPtrOrNil(strings.TrimSpace(sm.Route)),
				Icon:          strPtrOrNil(iconCopy),
				ParentID:      parentID,
				ApplicationID: &appIDCopy,
				SortOrder:     sm.SortOrder,
				Status:        status,
				CreatedAt:     now,
				UpdatedAt:     now,
			}

			if err := db.Create(&module).Error; err != nil {
				return fmt.Errorf("error creando módulo '%s' (route='%s', app='%s'): %w",
					sm.Name, sm.Route, appClientID, err)
			}
			logrus.Infof("Módulo creado: name='%s', route='%s', app='%s'", sm.Name, sm.Route, appClientID)
		} else {
			// 3.2) Actualizar módulo existente
			updates := map[string]any{
				"item":       &item,
				"name":       sm.Name,
				"sort_order": sm.SortOrder,
				"status":     status,
				"updated_at": now,
			}

			if sm.Route != "" {
				updates["route"] = sm.Route
			}
			if sm.Icon != "" {
				iconCopy := strings.TrimSpace(sm.Icon)
				updates["icon"] = &iconCopy
			}
			updates["parent_id"] = parentID

			if err := db.Model(&module).
				Where("id = ?", module.ID).
				Updates(updates).Error; err != nil {
				return fmt.Errorf("error actualizando módulo '%s' (route='%s', app='%s'): %w",
					sm.Name, sm.Route, appClientID, err)
			}
			logrus.Infof("Módulo actualizado: name='%s', route='%s', app='%s'", sm.Name, sm.Route, appClientID)
		}

		// 4) Crear ModuleRolePermission para cada rol listado
		for _, roleNameRaw := range sm.Roles {
			roleName := strings.TrimSpace(roleNameRaw)
			if roleName == "" {
				continue
			}

			// 4.1) Buscar ApplicationRole por nombre y application_id
			var appRole models.ApplicationRole
			if err := db.
				Where("LOWER(name) = LOWER(?) AND application_id = ?", roleName, app.ID).
				First(&appRole).Error; err != nil {
				return fmt.Errorf("no se encontró ApplicationRole name='%s' en app client_id='%s' para módulo '%s': %w",
					roleName, appClientID, sm.Name, err)
			}

			// 4.2) Verificar si ya existe ModuleRolePermission
			var mrp models.ModuleRolePermission
			err := db.
				Where("module_id = ? AND application_role_id = ?", module.ID, appRole.ID).
				First(&mrp).Error

			if err == nil {
				// Ya existe, no creamos de nuevo
				logrus.Infof("ModuleRolePermission ya existe: module='%s', role='%s', app='%s'",
					sm.Name, roleName, appClientID)
				continue
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("error verificando ModuleRolePermission módulo='%s', rol='%s': %w",
					sm.Name, roleName, err)
			}

			// 4.3) Crear nuevo ModuleRolePermission (permission_type = "access")
			mrp = models.ModuleRolePermission{
				ID:                uuid.New(),
				ModuleID:          module.ID,
				ApplicationRoleID: appRole.ID,
				PermissionType:    "access",
				CreatedAt:         now,
				IsDeleted:         false,
			}

			if err := db.Create(&mrp).Error; err != nil {
				return fmt.Errorf("error creando ModuleRolePermission módulo='%s', rol='%s', app='%s': %w",
					sm.Name, roleName, appClientID, err)
			}

			logrus.Infof("ModuleRolePermission creado: module='%s', role='%s', app='%s'",
				sm.Name, roleName, appClientID)
		}
	}

	return nil
}

// Helper para strings opcionales
func strPtrOrNil(s string) *string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	v := strings.TrimSpace(s)
	return &v
}
