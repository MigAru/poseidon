package upload

import (
	"context"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/sirupsen/logrus"
	"time"
)

type Manager struct {
	ctx            context.Context
	log            *logrus.Entry
	uploads        *Uploads
	defaultTimeout time.Duration
	chunkBus       chan Chunk
}

type Metadata struct {
	ID            string
	UploadedBytes int
}

func NewManager(ctx context.Context, cfg *config.Config, fs *file_system.FS, log *logrus.Logger) *Manager {
	return &Manager{
		ctx:            ctx,
		log:            log.WithField("prefix", "upload_manager"),
		uploads:        NewUploads(fs, log),
		defaultTimeout: time.Duration(cfg.Upload.MinutesTimeout) * time.Minute,
	}
}
