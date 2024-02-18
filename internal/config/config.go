package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB *DB
}

type DB struct {
	Name            string        `envconfig:"DATABASE_NAME" required:"true"`
	DriverName      string        `envconfig:"DATABASE_DRIVER_NAME" required:"true"`
	DSN             string        `envconfig:"DATABASE_DSN" required:"true"`
	MaxOpenConns    int           `envconfig:"DATABASE_MAX_OPEN_CONNS" required:"true"`
	MaxIdleConns    int           `envconfig:"DATABASE_MAX_IDLE_CONNS" required:"true"`
	ConnMaxLifetime time.Duration `envconfig:"DATABASE_CONN_MAX_LIFETIME" required:"true"`
}

func LoadConfig() *Config {
	for _, fileName := range []string{".env.local", ".env"} { //.env.local for local secrets (higher priority than .env)
		err := godotenv.Load(fileName) //in cycle cause first error in varargs prevents loading next files
		if err != nil {
			fmt.Println(fmt.Errorf("error loading %s fileName : %v", fileName, err))
		}
	}
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		fmt.Println(fmt.Errorf("cannot process envs: %v", err))
	} else {
		fmt.Println("Config initialized")
	}

	return &cfg
}
