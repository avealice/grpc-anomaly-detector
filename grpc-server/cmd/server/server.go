package main

import (
	"context"
	"log"

	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/app"
)

func main() {
	// grpcServer := grpc.NewServer()

	// pb.RegisterTransmitterServiceServer(grpcServer, &transmitter.TransmitterServer{})

	// port := ":3333"
	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }

	// log.Printf("Server listening on port %s\n", port)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
