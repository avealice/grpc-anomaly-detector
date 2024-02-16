package detector

func (i *Implementation) Connect(id string, mean, std, k float64) error {
	err := i.detectorService.Connect(id, mean, std, k)
	if err != nil {
		return err
	}

	return nil
}
