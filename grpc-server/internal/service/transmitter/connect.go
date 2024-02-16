package transmitter

import (
	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/model"
)

func (s *service) Connect() model.Connection {
	return model.Connection{
		UUID: s.uuid,
		Mean: s.mean,
		Std:  s.std,
	}
}
