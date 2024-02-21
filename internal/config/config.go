package config

import (
	"github.com/plitn/wh_api_lamoda/internal/logger"
	"os"
)

type Config struct {
	DB *DB
}

type DB struct {
	DSN        string
	DriverName string
}

func LoadConfig() *Config {
	logger.Logger.Println("1")

	driverName := os.Getenv("DATABASE_DRIVER_NAME")
	logger.Logger.Println("2")

	if driverName == "" {
		logger.Logger.Println("DATABASE_DRIVER_NAME is not set")
	}
	logger.Logger.Println("3")

	dsn := os.Getenv("DATABASE_DSN")
	logger.Logger.Println("4")

	if dsn == "" {
		logger.Logger.Println("DATABASE_DSN is not set")
	}
	logger.Logger.Println("5")

	cfg := Config{
		DB: &DB{
			DSN:        dsn,
			DriverName: driverName,
		},
	}
	logger.Logger.Println("6")

	return &cfg
}
