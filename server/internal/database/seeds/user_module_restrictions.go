package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"central-user-manager/internal/config"
	"central-user-manager/internal/models"
)

type SeedUserModuleRestriction struct {
	UserEmail           string  `json:"user_email"`
	ApplicationClientID string  `json:"application_client_id,omitempty"` // preferido si viene
	ApplicationName     string  `json:"application_name,omitempty"`      // fallback si no viene client_id
	ModuleName          string  `json:"module_name"`
	ModuleRoute         string  `json:"module_route"`
	RestrictionType     string  `json:"restriction_type"`               // ej: "limit_permission"
	MaxPermissionLevel  *string `json:"max_permission_level,omitempty"` // requerido si restriction_type="limit_permission" ("denied" | "access")
	Reason              *string `json:"reason,omitempty"`
	ExpiresAt           *string `json:"expires_at,omitempty"` // RFC3339 o vacío
	CreatedByEmail      string  `json:"created_by_email"`
}

func SeedUserModuleRestrictions() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding user_module_restrictions desde JSON...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	entries, err := loadUMRJSON("internal/data/user_module_restrictions.json")
	if err != nil {
		return err
	}

	for _, e := range entries {
		if err := config.DB.Transaction(func(tx *gorm.DB) error {
			return upsertOneUMRTx(tx, e)
		}); err != nil {
			return err
		}
	}

	logrus.Info("Seeding user_module_restrictions: COMPLETADO")
	return nil
}

func loadUMRJSON(path string) ([]SeedUserModuleRestriction, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir %s: %w", path, err)
	}
	defer func() { _ = f.Close() }()

	var input []SeedUserModuleRestriction
	if err := json.NewDecoder(f).Decode(&input); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON de user_module_restrictions: %w", err)
	}
	return input, nil
}

func upsertOneUMRTx(tx *gorm.DB, e SeedUserModuleRestriction) error {
	var user models.User
	if err := tx.Where("LOWER(email) = LOWER(?)", e.UserEmail).First(&user).Error; err != nil {
		return fmt.Errorf("no se encontró user_email='%s': %w", e.UserEmail, err)
	}

	var createdBy models.User
	if err := tx.Where("LOWER(email) = LOWER(?)", e.CreatedByEmail).First(&createdBy).Error; err != nil {
		createdBy = user
		logrus.Warnf("No se encontró created_by '%s', usando el mismo user '%s'", e.CreatedByEmail, e.UserEmail)
	}

	var app models.Application
	if strings.TrimSpace(e.ApplicationClientID) != "" {
		if err := tx.Where("client_id = ?", e.ApplicationClientID).First(&app).Error; err != nil {
			return fmt.Errorf("no se encontró application client_id='%s': %w", e.ApplicationClientID, err)
		}
	} else {
		if err := tx.Where("LOWER(name) = LOWER(?)", e.ApplicationName).First(&app).Error; err != nil {
			return fmt.Errorf("no se encontró application name='%s': %w", e.ApplicationName, err)
		}
	}

	var module models.Module
	q := tx.Where("application_id = ? AND route = ?", app.ID, e.ModuleRoute).First(&module)
	if q.Error != nil {
		if err := tx.Where("application_id = ? AND LOWER(name) = LOWER(?)", app.ID, e.ModuleName).
			First(&module).Error; err != nil {
			return fmt.Errorf("no se encontró módulo (name='%s', route='%s') en app '%s': %w",
				e.ModuleName, e.ModuleRoute, app.Name, err)
		}
	}

	now := time.Now()
	var exp *time.Time
	if e.ExpiresAt != nil && strings.TrimSpace(*e.ExpiresAt) != "" {
		t, perr := time.Parse(time.RFC3339, strings.TrimSpace(*e.ExpiresAt))
		if perr != nil {
			return fmt.Errorf("expires_at inválido '%s': %w", *e.ExpiresAt, perr)
		}
		exp = &t
	}

	if strings.EqualFold(e.RestrictionType, "limit_permission") {
		if e.MaxPermissionLevel == nil || (strings.ToLower(*e.MaxPermissionLevel) != "denied" && strings.ToLower(*e.MaxPermissionLevel) != "access") {
			return fmt.Errorf("max_permission_level requerido ('denied'|'access') para restriction_type='limit_permission'")
		}
	}

	var count int64
	if err := tx.Model(&models.UserModuleRestriction{}).
		Where("user_id = ? AND module_id = ? AND application_id = ? AND is_deleted = false",
			user.ID, module.ID, app.ID).
		Count(&count).Error; err != nil {
		return fmt.Errorf("error verificando duplicado de UMR user='%s' modulo='%s' app='%s': %w",
			user.Email, module.Name, app.Name, err)
	}

	if count > 0 {
		upd := map[string]any{
			"restriction_type":     e.RestrictionType,
			"max_permission_level": e.MaxPermissionLevel,
			"reason":               e.Reason,
			"expires_at":           exp,
			"updated_at":           now,
			"updated_by":           createdBy.ID,
		}
		if err := tx.Model(&models.UserModuleRestriction{}).
			Where("user_id = ? AND module_id = ? AND application_id = ? AND is_deleted = false",
				user.ID, module.ID, app.ID).
			Updates(upd).Error; err != nil {
			return fmt.Errorf("error actualizando UMR user='%s' modulo='%s' app='%s': %w",
				user.Email, module.Name, app.Name, err)
		}
		logrus.Warnf("UMR actualizado (ya existía): user='%s' modulo='%s' app='%s'", user.Email, module.Name, app.Name)
		return nil
	}

	umr := models.UserModuleRestriction{
		ID:                 uuid.New(),
		UserID:             user.ID,
		ModuleID:           module.ID,
		ApplicationID:      app.ID,
		RestrictionType:    e.RestrictionType,
		MaxPermissionLevel: e.MaxPermissionLevel,
		Reason:             e.Reason,
		ExpiresAt:          exp,
		CreatedAt:          now,
		CreatedBy:          createdBy.ID,
		UpdatedAt:          now,
		UpdatedBy:          &createdBy.ID,
		IsDeleted:          false,
	}

	if err := tx.Create(&umr).Error; err != nil {
		return fmt.Errorf("error insertando UMR user='%s' modulo='%s' app='%s': %w",
			user.Email, module.Name, app.Name, err)
	}

	logrus.Infof("UMR insertado: user='%s' modulo='%s' app='%s'", user.Email, module.Name, app.Name)
	return nil
}
