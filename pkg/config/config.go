package config

import (
	"time"

	"github.com/joho/godotenv"
)

// GRPCConfig ...
type GRPCConfig interface {
	Address() string
}

// PGConfig ...
type PGConfig interface {
	DSN() string
}

// HTTPConfig ...
type HTTPConfig interface {
	Address() string
}

// SwaggerConfig ...
type SwaggerConfig interface {
	Address() string
}

// RedisConfig ...
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

// StorageConfig ...
type StorageConfig interface {
	Mode() string
}

// Load ...
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
