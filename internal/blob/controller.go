package blob

import (
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/sirupsen/logrus"
	httpInterface "net/http"
)

type Controller struct {
	log       *logrus.Logger
	fs        *file_system.FS
	chunkSize int
	manager   *upload.Manager
}

func NewController(log *logrus.Logger, cfg *config.Config, fs *file_system.FS, manager *upload.Manager) *Controller {
	return &Controller{log: log, chunkSize: cfg.Upload.ChunkSize, fs: fs, manager: manager}
}

func (c *Controller) buildStatusUpload(uploadedSize, totalSize int) int {
	if uploadedSize == totalSize {
		return httpInterface.StatusCreated
	}
	return httpInterface.StatusAccepted
}
