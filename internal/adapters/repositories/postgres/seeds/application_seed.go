package seeds

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
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
	file, err := os.Open("internal/adapters/repositories/data/applications.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer file.Close()

	var apps []SeedApplication
	if err := json.NewDecoder(file).Decode(&apps); err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	for _, a := range apps {
		// Verificar si el ClientID ya existe
		var exists bool
		err := database.DB.Model(&domain.Application{}).
			Select("count(*) > 0").
			Where("client_id = ?", a.ClientID).
			Find(&exists).Error
		if err != nil {
			return fmt.Errorf("error al verificar existencia de client_id '%s': %w", a.ClientID, err)
		}
		if exists {
			continue
		}

		app := domain.Application{
			ID:           uuid.New(),
			Name:         a.Name,
			ClientID:     a.ClientID,
			ClientSecret: a.ClientSecret,
			Domain:       a.Domain,
			Logo:         &a.Logo,
			Description:  &a.Description,
			CallbackUrls: pq.StringArray(a.CallbackURLs),
			IsFirstParty: a.IsFirstParty,
			Status:       a.Status,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			IsDeleted:    false,
		}

		if err := database.DB.Create(&app).Error; err != nil {
			return fmt.Errorf("error al insertar aplicación '%s': %w", a.Name, err)
		}
	}

	return nil
}
