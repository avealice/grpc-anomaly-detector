package detector

import (
	desc "github.com/avealice/grpc-anomaly-detector/grpc-server/pkg/transmitter"
)

func (i *Implementation) StreamData(msg *desc.Message) {
	i.detectorService.StreamData(msg)
}
