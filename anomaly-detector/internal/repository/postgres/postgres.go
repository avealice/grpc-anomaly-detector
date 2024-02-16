package postgres

import (
	"fmt"

	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/repository/postgres/model"

	"github.com/avealice/grpc-anomaly-detector/anomaly-detector/internal/repository/postgres/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	config config.Config
	db     *gorm.DB
}

func NewPostrgesDB() *postgresDB {
	return &postgresDB{
		config: *config.NewConfig(),
	}
}

func (p *postgresDB) Connect() error {
	dsn := fmt.Sprintf("user=%s dbname=%s port=%s sslmode=%s", p.config.UserName, p.config.DBName, p.config.Port, p.config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	p.db = db

	if !p.db.Migrator().HasTable(&model.Anomaly{}) {
		err := p.createAnomalyTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *postgresDB) createAnomalyTable() error {
	err := p.db.AutoMigrate(&model.Anomaly{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate Anomaly table: %v", err)
	}

	return nil
}

func (p *postgresDB) Disconnect() error {
	if p.db != nil {
		db, err := p.db.DB()
		if err != nil {
			return err
		}
		db.Close()
	}

	return nil
}

func (p *postgresDB) CreateAnomaly(anomaly *model.Anomaly) error {
	return p.db.Create(anomaly).Error
}
