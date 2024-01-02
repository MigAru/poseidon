package blob

import (
	"github.com/MigAru/poseidon/internal/database"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/sirupsen/logrus"
	httpInterface "net/http"
)

type Controller struct {
	log     *logrus.Logger
	fs      *file_system.FS
	uploads *upload.Uploads
	db      database.DB
}

func NewController(log *logrus.Logger, fs *file_system.FS, uploads *upload.Uploads) *Controller {
	return &Controller{log: log, fs: fs, uploads: uploads}
}

func (c *Controller) buildStatusUpload(uploadedSize, totalSize int) int {
	if uploadedSize == totalSize {
		return httpInterface.StatusCreated
	}
	return httpInterface.StatusAccepted
}
