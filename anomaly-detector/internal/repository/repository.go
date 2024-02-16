package repository

import "github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/repository/postgres/model"

type DetectorRepository interface {
	Connect() error
	Disconnect() error
	CreateAnomaly(anomaly *model.Anomaly) error
}
