package seeds

import (
	"fmt"
	"os"
	"time"

	"server/internal/models"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type SeedUbigeoData struct {
	Ubigeos []SeedUbigeo `yaml:"ubigeos"`
}

type SeedUbigeo struct {
	Ubdep   string `yaml:"ubdep"`
	Ubprv   string `yaml:"ubprv"`
	Ubdis   string `yaml:"ubdis"`
	Nodep   string `yaml:"nodep"`
	Noprv   string `yaml:"noprv"`
	Nodis   string `yaml:"nodis"`
	Cpdis   string `yaml:"cpdis"`
	UbInei  string `yaml:"ub_inei"`
}

func SeedUbigeos(db *gorm.DB) error {
	logrus.Info("----------------------------------------------------------------------------------------------")
	logrus.Info("Seeding ubigeos desde YAML...")
	logrus.Info("----------------------------------------------------------------------------------------------")

	filePath := "internal/database/seeds/data/ubigeos.yml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo YAML (%s): %w", filePath, err)
	}

	var ubigeoData SeedUbigeoData
	if err := yaml.Unmarshal(data, &ubigeoData); err != nil {
		return fmt.Errorf("error al decodificar YAML: %w", err)
	}

	for _, u := range ubigeoData.Ubigeos {
		// Construir el código de ubigeo concatenando ubdep + ubprv + ubdis
		ubigeoCode := u.Ubdep + u.Ubprv + u.Ubdis

		var count int64
		err := db.Model(&models.Ubigeo{}).
			Where("ubigeo_code = ?", ubigeoCode).
			Count(&count).Error
		if err != nil {
			return fmt.Errorf("error al verificar existencia de ubigeo '%s': %w", ubigeoCode, err)
		}

		if count > 0 {
			logrus.Warnf("Ubigeo ya existe: %s (%s - %s - %s)", ubigeoCode, u.Nodep, u.Noprv, u.Nodis)
			continue
		}

		ubigeo := models.Ubigeo{
			UbigeoCode: ubigeoCode,
			IneiCode:   u.UbInei,
			Department: u.Nodep,
			Province:   u.Noprv,
			District:   u.Nodis,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := db.Create(&ubigeo).Error; err != nil {
			return fmt.Errorf("error al insertar ubigeo '%s': %w", ubigeoCode, err)
		}

		logrus.Infof("Ubigeo insertado: %s (%s - %s - %s)", ubigeoCode, u.Nodep, u.Noprv, u.Nodis)
	}

	return nil
}