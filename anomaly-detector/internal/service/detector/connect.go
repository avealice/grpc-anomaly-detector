package detector

func (s *service) Connect(id string, mean, std, k float64) error {
	err := s.detectorRepository.Connect()
	if err != nil {
		return err
	}

	defer s.detectorRepository.Disconnect()

	// fmt.Printf("id %s, mean %f, std %f, k %f", id, mean, std, k)
	s.detectorStats.AnomalyCoefficient = k
	s.detectorStats.Id = id
	s.detectorStats.Mean = mean
	s.detectorStats.Std = std

	return nil
}
