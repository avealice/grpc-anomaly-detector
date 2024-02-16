package detector

import (
	"log"
	"math"

	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/model"

	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/repository"
)

type service struct {
	detectorRepository repository.DetectorRepository
	detectorStats      model.DetectorStats
}

func NewService(detectorRepository repository.DetectorRepository) *service {
	return &service{
		detectorRepository: detectorRepository,
	}
}

func (s *service) UpdateStats(frequency float64) {
	s.detectorStats.ValuesProcessed++

	delta := frequency - s.detectorStats.Mean
	s.detectorStats.Mean += delta / float64(s.detectorStats.ValuesProcessed)
	delta2 := frequency - s.detectorStats.Mean
	s.detectorStats.SumSquares += delta * delta2

	if s.detectorStats.ValuesProcessed > 30 && !s.detectorStats.DetectionMode {
		s.detectorStats.Std = math.Sqrt(s.detectorStats.SumSquares / float64(s.detectorStats.ValuesProcessed-1))

		log.Printf("Approximation - Count: %d, Mean: %f, STD: %f\n", s.detectorStats.ValuesProcessed, s.detectorStats.Mean, s.detectorStats.Std)

		if s.detectorStats.Std != 0 {
			s.detectorStats.DetectionMode = true
		}
	}

}

func (s *service) CheckAnomalies(frequency float64) {
	if s.detectorStats.DetectionMode && s.isAnomaly(frequency) {
		s.detectorStats.AnomalyDetected = true
	}

}

func (s *service) isAnomaly(frequency float64) bool {
	return math.Abs(frequency-s.detectorStats.Mean) > s.detectorStats.AnomalyCoefficient*s.detectorStats.Std
}
