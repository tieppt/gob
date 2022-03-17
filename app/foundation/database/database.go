package database

import (
	"fmt"
	"os"

	"github.com/tieppt/gob/app/foundation/config"
	"github.com/tieppt/gob/app/foundation/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	DB *gorm.DB
}

var DBInst DBConn

func ConnectDB(logger *zap.Logger) {
	// Get DB connection string from environment variables
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Conf.PGHostname,
		config.Conf.PGPort,
		config.Conf.PGUser,
		config.Conf.PGDatabase,
		config.Conf.PGPassword,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Fatal("Connect to database failed", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("Connected to database")
	DBInst.DB = db

	// Migrate the schema
	if err := DBInst.DB.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		logger.Fatal("Migrate database failed", zap.Error(err))
		os.Exit(2)
	}
	logger.Info("Migrated database")
}
