package transmitter

import (
	"math/rand"

	"github.com/google/uuid"

	def "github.com/avealice/grpc-anomaly-detector/grpc-server/internal/service"
)

var _ def.TransmitterService = (*service)(nil)

type service struct {
	uuid string
	mean float64
	std  float64
}

func NewService() *service {
	return &service{
		uuid: newUUID(),
		mean: newMean(),
		std:  newSTD(),
	}
}

func newUUID() string {
	return uuid.New().String()
}

func newMean() float64 {
	return rand.Float64()*20 - 10 // [-10, 10]
}

func newSTD() float64 {
	return rand.Float64()*1.2 + 0.3 // [0.3, 1.5]
}
