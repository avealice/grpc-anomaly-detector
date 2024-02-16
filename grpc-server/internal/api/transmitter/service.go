package transmitter

import (
	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/service"
)

type Implementation struct {
	transmitterService service.TransmitterService
}

func NewImplementation(transmitterService service.TransmitterService) *Implementation {
	return &Implementation{
		transmitterService: transmitterService,
	}
}
