package seeds

func Run() error {
	if err := SeedStructuralPositions(); err != nil {
		return err
	}

	if err := SeedOrganicUnits(); err != nil {
		return err
	}

	return nil
}
