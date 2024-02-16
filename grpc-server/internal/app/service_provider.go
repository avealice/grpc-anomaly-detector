package app

import (
	"log"

	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/config"

	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/api/transmitter"
	"github.com/avealice/grpc-anomaly-detector/grpc-server/internal/service"
	transmitterService "github.com/avealice/grpc-anomaly-detector/grpc-server/internal/service/transmitter"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	transmitterService service.TransmitterService

	transmitterImpl *transmitter.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) TransmitterService() service.TransmitterService {
	if s.transmitterService == nil {
		s.transmitterService = transmitterService.NewService()
	}

	return s.transmitterService
}

func (s *serviceProvider) TransmitterImpl() *transmitter.Implementation {
	if s.transmitterImpl == nil {
		s.transmitterImpl = transmitter.NewImplementation(s.TransmitterService())
	}

	return s.transmitterImpl
}
