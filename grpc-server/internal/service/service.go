package service

import (
	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/model"
)

type TransmitterService interface {
	Connect() model.Connection
	StreamData() model.StreamData
}
