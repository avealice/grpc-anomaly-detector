package transmitter

import (
	"math/rand"
	"time"

	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/model"
	"gonum.org/v1/gonum/stat/distuv"
)

func (s *service) StreamData() model.StreamData {
	return model.StreamData{
		SessionID: s.uuid,
		Frequency: s.newFrequency(),
		Timestamp: time.Now().Unix(),
	}
}

func (s *service) newFrequency() float64 {
	normalDist := distuv.Normal{
		Mu:    s.mean,
		Sigma: s.std,
	}

	rand.Seed(time.Now().UnixNano())
	return normalDist.Rand()

}
