package file_system

import (
	"github.com/MigAru/poseidon/internal/config"
	"time"
)

type FS struct {
	basePath   string
	blobExpire time.Duration
}

func New(cfg *config.Config) *FS {
	return &FS{basePath: cfg.FileSystem.BasePath, blobExpire: time.Second}
}
