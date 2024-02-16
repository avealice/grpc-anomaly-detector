package app

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/avealice/grpc-anomaly-detector/grpc-server/pkg/transmitter"
)

type App struct {
	serviceProvider *serviceProvider
	grpcClient      desc.TransmitterServiceClient
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runGRPCClient()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	a.serviceProvider.DetectorImpl()

	return nil
}

// func (a *App) initGRPCClient(_ context.Context) error {

// }

func (a *App) runGRPCClient() error {
	k := flag.Float64("k", 0.0, "Coefficient for anomaly detection")
	flag.Parse()

	conn, err := grpc.Dial(a.serviceProvider.GPRSConfig().Address(), grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()

	log.Printf("GRPC client is running on %s", a.serviceProvider.GPRSConfig().Address())

	a.grpcClient = desc.NewTransmitterServiceClient(conn)

	response, err := a.grpcClient.Connect(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling Connect: %v", err)
	}

	a.serviceProvider.detectorImpl.Connect(response.SessionId, response.Mean, response.Std, *k)

	stream, err := a.grpcClient.StreamData(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling StreamData: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
			break
		}

		a.serviceProvider.detectorImpl.StreamData(message)

		// time.Sleep(time.Millisecond * 500)
	}

	return nil
}
