package seeds

import (
	"server/pkg/logger"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	seeders := []func(*gorm.DB) error{
		SeedStructuralPositions,
		SeedOrganicUnits,
		SeedApplications,
		SeedApplicationRoles,
		SeedModules,
		SeedUsersAndUserApplicationRoles,
	}

	for _, seed := range seeders {
		if err := seed(db); err != nil {
			return err
		}
	}

	logger.Log.Info("Todos los seeders se ejecutaron correctamente")
	return nil
}
