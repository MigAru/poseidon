package upload

import (
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/pkg/registry/hasher"
	"github.com/sirupsen/logrus"
)

type Uploads struct {
	fs     *file_system.FS
	log    *logrus.Logger
	hasher *hasher.Hasher
}

func NewUploads(fs *file_system.FS, log *logrus.Logger, hasher *hasher.Hasher) *Uploads {
	return &Uploads{
		hasher: hasher,
		fs:     fs,
		log:    log,
	}
}
