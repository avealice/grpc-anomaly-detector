package detector

import (
	"fmt"
	"log"

	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/repository/postgres/model"

	desc "github.com/avealice/grpc-anomaly-detector/grpc-server/pkg/transmitter"
)

func (s *service) StreamData(msg *desc.Message) {
	s.UpdateStats(msg.Frequency)

	if s.detectorStats.ValuesProcessed%30 == 0 {
		log.Printf("Processed %d values. Current Mean: %.2f, Current Std: %.2f",
			s.detectorStats.ValuesProcessed, s.detectorStats.Mean, s.detectorStats.Std)
	}

	s.CheckAnomalies(msg.Frequency)

	if s.detectorStats.AnomalyDetected {
		log.Printf("Anomaly detected: Frequency=%.2f, Mean=%.2f, Std=%.2f",
			msg.Frequency, s.detectorStats.Mean, s.detectorStats.Std)

		anomaly := &model.Anomaly{
			SessionID:    msg.SessionId,
			Frequency:    msg.Frequency,
			TimestampUTC: msg.TimestampUtc,
		}

		s.detectorRepository.CreateAnomaly(anomaly)
	} else {
		fmt.Printf("Session ID: %s, Frequency: %f, Timestamp UTC: %d\n",
			msg.SessionId, msg.Frequency, msg.TimestampUtc)
	}
}
