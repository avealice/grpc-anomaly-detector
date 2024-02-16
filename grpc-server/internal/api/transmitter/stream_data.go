package transmitter

import (
	pb "github.com/avealice/grpc-anomaly-detector/grpc-server/pkg/transmitter"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) StreamData(empty *emptypb.Empty, stream pb.TransmitterService_StreamDataServer) error {
	for {
		streamData := i.transmitterService.StreamData()

		if err := stream.Send(&pb.Message{
			SessionId:    streamData.SessionID,
			Frequency:    streamData.Frequency,
			TimestampUtc: streamData.Timestamp}); err != nil {
			return err
		}
	}
}
