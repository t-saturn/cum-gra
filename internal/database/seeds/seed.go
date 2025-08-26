// Package seeds contiene funciones para poblar la base de datos con datos iniciales o de prueba.
package seeds

// Run ejecuta todas las funciones de seeding para poblar la base de datos con datos iniciales.
func Run() error {
	if err := SeedStructuralPositions(); err != nil {
		return err
	}

	if err := SeedOrganicUnits(); err != nil {
		return err
	}

	if err := SeedApplications(); err != nil {
		return err
	}

	if err := SeedModules(); err != nil {
		return err
	}

	if err := SeedApplicationRoles(); err != nil {
		return err
	}

	if err := SeedUsersAndUserApplicationRoles(); err != nil {
		return err
	}

	if err := SeedUserModuleRestrictions(); err != nil {
		return err
	}

	return nil
}
