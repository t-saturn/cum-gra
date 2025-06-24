package seeds

func Run() error {
	if err := SeedStructuralPositions(); err != nil {
		return err
	}

	return nil
}
