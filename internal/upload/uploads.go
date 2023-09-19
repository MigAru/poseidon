package upload

import (
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/sirupsen/logrus"
	"sync"
)

type Uploads struct {
	mu     sync.RWMutex
	fs     *file_system.FS
	log    *logrus.Logger
	unsafe map[string]*Upload
}

func NewUploads(fs *file_system.FS, log *logrus.Logger) *Uploads {
	return &Uploads{
		unsafe: make(map[string]*Upload),
		fs:     fs,
		log:    log,
	}
}
