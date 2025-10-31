package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SeedApplication struct {
	Name         string   `json:"name"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Domain       string   `json:"domain"`
	Logo         string   `json:"logo"`
	Description  string   `json:"description"`
	CallbackURLs []string `json:"callback_urls"`
	IsFirstParty bool     `json:"is_first_party"`
	Status       string   `json:"status"`
}

func SeedApplications() error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding aplicaciones desde JSON...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	file, err := os.Open("internal/data/applications.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error al cerrar el archivo: %v\n", cerr)
		}
	}()

	var apps []SeedApplication
	if err := json.NewDecoder(file).Decode(&apps); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, a := range apps {
		var count int64
		if err := config.DB.Model(&models.Application{}).
			Where("client_id = ?", a.ClientID).
			Count(&count).Error; err != nil {
			return fmt.Errorf("error al verificar existencia de client_id '%s': %w", a.ClientID, err)
		}

		if count > 0 {
			logrus.Warnf("Aplicación ya existe: %s", a.Name)
			continue
		}

		app := models.Application{
			ID:           uuid.New(),
			Name:         a.Name,
			ClientID:     a.ClientID,
			ClientSecret: a.ClientSecret,
			Domain:       a.Domain,
			Logo:         &a.Logo,
			Description:  &a.Description,
			Status:       a.Status,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			IsDeleted:    false,
		}

		if err := config.DB.Create(&app).Error; err != nil {
			return fmt.Errorf("error al insertar aplicación '%s': %w", a.Name, err)
		}

		logrus.Infof("Aplicación insertada: %s", a.Name)
	}

	return nil
}
