package seeds

func Run() error {
	seeders := []func() error{
		SeedStructuralPositions,
		SeedOrganicUnits,
		SeedApplications,
		SeedModules,
		SeedApplicationRoles,
		SeedUsersAndUserApplicationRoles,
		SeedUserModuleRestrictions,
	}

	for _, seed := range seeders {
		if err := seed(); err != nil {
			return err
		}
	}

	return nil
}
