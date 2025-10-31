package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SeedModuleRolePermission struct {
	ApplicationName     string   `json:"application_name"`
	ApplicationRoleName string   `json:"application_role_name"`
	PermissionType      string   `json:"permission_type"`
	Modules             []string `json:"modules"`
}

func SeedModuleRolePermissions() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding permisos de roles por módulo desde JSON...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	file, err := os.Open("internal/data/module_role_permissions.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar el archivo: %v\n", cerr)
		}
	}()

	var seedData []SeedModuleRolePermission
	if err := json.NewDecoder(file).Decode(&seedData); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	db := config.DB
	insertedCount := 0

	for _, s := range seedData {
		var app models.Application
		if err := db.Where("LOWER(name) = ?", strings.ToLower(s.ApplicationName)).First(&app).Error; err != nil {
			return fmt.Errorf("no se encontró la aplicación '%s': %w", s.ApplicationName, err)
		}

		var role models.ApplicationRole
		if err := db.Where("LOWER(name) = ? AND application_id = ?", strings.ToLower(s.ApplicationRoleName), app.ID).First(&role).Error; err != nil {
			return fmt.Errorf("no se encontró el rol '%s' para la aplicación '%s': %w", s.ApplicationRoleName, s.ApplicationName, err)
		}

		for _, moduleName := range s.Modules {
			var module models.Module
			if err := db.Where("LOWER(name) = ? AND application_id = ?", strings.ToLower(moduleName), app.ID).First(&module).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					logrus.Warnf("No se encontró el módulo '%s' en la aplicación '%s'. Se omite.", moduleName, s.ApplicationName)
					continue
				}
				return fmt.Errorf("error al buscar módulo '%s': %w", moduleName, err)
			}

			var count int64
			if err := db.Model(&models.ModuleRolePermission{}).
				Where("module_id = ? AND application_role_id = ? AND permission_type = ?", module.ID, role.ID, s.PermissionType).
				Count(&count).Error; err != nil {
				return fmt.Errorf("error al verificar duplicados: %w", err)
			}
			if count > 0 {
				logrus.Warnf("Permiso ya existente: módulo '%s', rol '%s', app '%s'", moduleName, s.ApplicationRoleName, s.ApplicationName)
				continue
			}

			mrp := models.ModuleRolePermission{
				ID:                uuid.New(),
				ModuleID:          module.ID,
				ApplicationRoleID: role.ID,
				PermissionType:    s.PermissionType,
				CreatedAt:         time.Now(),
				IsDeleted:         false,
			}

			if err := db.Create(&mrp).Error; err != nil {
				return fmt.Errorf("error al insertar permiso para módulo '%s': %w", moduleName, err)
			}

			insertedCount++
			logrus.Infof("Permiso insertado: [%s] módulo '%s' → rol '%s'", s.ApplicationName, moduleName, s.ApplicationRoleName)
		}
	}

	logrus.Infof("Total de permisos insertados: %d", insertedCount)
	logrus.Info("----------------------------------------------------------------------------------------------")
	return nil
}
