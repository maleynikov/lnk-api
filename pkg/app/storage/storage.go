package storage

import (
	"fmt"
	"time"

	"gitlab.maleynikov.me/url-short/api/pkg/app"
	"gitlab.maleynikov.me/url-short/api/pkg/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	DB *gorm.DB
}

// func (s *Storage) Close() {
// 	db, _ := s.DB.DB()
// 	db.Close()
// }

func (s *Storage) Migrate() {
	s.DB.AutoMigrate(&models.User{})
	s.DB.AutoMigrate(&models.Url{})
	s.DB.AutoMigrate(&models.Statistic{})
}

var storage *Storage

func NewStorage(cfg *app.Config) (*Storage, error) {
	if storage != nil {
		return storage, nil
	}
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("conn error %w", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	storage := &Storage{
		DB: db,
	}
	// storage.Migrate()

	return storage, nil
}
