package env

import (
	"os"

	"github.com/pkg/errors"
)

const storageModeEnvName = "STORAGE_MODE"

// StorageConfig ...
type StorageConfig struct {
	mode string
}

// NewStorageConfig ...
func NewStorageConfig() (*StorageConfig, error) {
	storageMode := os.Getenv(storageModeEnvName)
	if len(storageMode) == 0 {
		return nil, errors.New("storage mode not found")
	}

	return &StorageConfig{
		mode: storageMode,
	}, nil
}

// Mode ...
func (cfg *StorageConfig) Mode() string {
	return cfg.mode
}
