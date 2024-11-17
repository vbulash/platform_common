package env

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "DB_DSN"
)

// PGConfig ...
type PGConfig struct {
	dsn string
}

// NewPGConfig ...
func NewPGConfig() (*PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &PGConfig{
		dsn: dsn,
	}, nil
}

// DSN ...
func (cfg *PGConfig) DSN() string {
	return cfg.dsn
}
