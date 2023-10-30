package upload

import (
	"context"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/pkg/registry/hasher"
	"github.com/sirupsen/logrus"
)

type Manager struct {
	ctx     context.Context
	log     *logrus.Entry
	hasher  *hasher.Hasher
	uploads *Uploads
}

func NewManager(ctx context.Context, fs *file_system.FS, hasher *hasher.Hasher, log *logrus.Logger) *Manager {
	return &Manager{
		ctx:     ctx,
		log:     log.WithField("prefix", "upload_manager"),
		uploads: NewUploads(fs, log),
		hasher:  hasher,
	}
}
