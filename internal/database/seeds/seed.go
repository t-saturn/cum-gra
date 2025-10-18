package seeds

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
