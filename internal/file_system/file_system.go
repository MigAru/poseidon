package file_system

import "github.com/MigAru/poseidon/internal/config"

type FS struct {
	basePath string
}

func New(cfg *config.Config) *FS {
	return &FS{basePath: cfg.FileSystem.BasePath}
}
