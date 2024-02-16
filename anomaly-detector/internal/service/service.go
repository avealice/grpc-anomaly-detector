package service

import (
	desc "github.com/avealice/grpc-anomaly-detector/grpc-server/pkg/transmitter"
)

type DetectorService interface {
	Connect(id string, mean, std, k float64) error
	StreamData(msg *desc.Message)
}
