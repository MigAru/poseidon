package manifest

import (
	"github.com/MigAru/poseidon/internal/database"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/MigAru/poseidon/pkg/registry/hasher"
	"github.com/sirupsen/logrus"
	"strings"
)

type Controller struct {
	log     *logrus.Logger
	fs      *file_system.FS
	hr      hasher.Hasher
	db      database.DB
	uploads *upload.Manager
}

//TODO: сделать обработку ошибок
//TODO: разнести функции

func NewController(log *logrus.Logger, uploads *upload.Manager, fs *file_system.FS, db database.DB) *Controller {
	return &Controller{
		log:     log,
		fs:      fs,
		db:      db,
		uploads: uploads,
	}
}

func (c *Controller) isDigest(name string) bool {
	hashArray := strings.Split(name, ":")
	return len(hashArray) > 1
}
