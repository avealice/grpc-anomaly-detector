package detector

import (
	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/service"
)

type Implementation struct {
	detectorService service.DetectorService
}

func NewImplementation(detectorService service.DetectorService) *Implementation {
	return &Implementation{
		detectorService: detectorService,
	}
}
