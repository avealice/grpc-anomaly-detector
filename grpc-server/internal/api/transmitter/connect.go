package transmitter

import (
	"context"
	"log"

	pb "github.com/avealice/grpc-anomaly-detector/grpc-server/pkg/transmitter"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Connect(ctx context.Context, empty *emptypb.Empty) (*pb.NewConnection, error) {
	connect := i.transmitterService.Connect()

	log.Printf("New connection - SessionID: %s, Mean: %f, STD: %f\n", connect.UUID, connect.Mean, connect.Std)

	return &pb.NewConnection{
		SessionId: connect.UUID,
		Mean:      connect.Mean,
		Std:       connect.Std,
	}, nil
}
