package app

import (
	"log"

	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/api/detector"
	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/config"
	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/repository"
	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/repository/postgres"
	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/service"
	detectorService "github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/service/detector"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	detectorService service.DetectorService

	detectorRepository repository.DetectorRepository

	detectorImpl *detector.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GPRSConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DetectorRepository() repository.DetectorRepository {
	if s.detectorRepository == nil {
		s.detectorRepository = postgres.NewPostrgesDB()
	}

	return s.detectorRepository
}

func (s *serviceProvider) DetectorService() service.DetectorService {
	if s.detectorService == nil {
		s.detectorService = detectorService.NewService(s.DetectorRepository())
	}

	return s.detectorService
}

func (s *serviceProvider) DetectorImpl() *detector.Implementation {
	if s.detectorImpl == nil {
		s.detectorImpl = detector.NewImplementation(s.DetectorService())
	}

	return s.detectorImpl
}
