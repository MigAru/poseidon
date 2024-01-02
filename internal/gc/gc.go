package gc

import (
	"github.com/MigAru/poseidon/internal/database"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type GC struct {
	log *logrus.Entry
	fs  *file_system.FS
	db  database.DB
}

func New(log *logrus.Logger, fs *file_system.FS, db database.DB) *GC {
	return &GC{
		fs:  fs,
		db:  db,
		log: log.WithField("prefix", "garbage_collector"),
	}
}

func (gc *GC) Clear(ctx http.Context) {
	go gc.clear(ctx)
}

func (gc *GC) clear(ctx http.Context) {
	defer func() func() {
		startTime := time.Now()
		return func() {
			gc.log.
				WithField("operation", "gc_execution_time").
				Info(time.Since(startTime).String())
		}
	}()()

	full, _ := strconv.ParseBool(ctx.QueryParam("full"))

	if err := gc.clearBlobs(); err != nil {
		gc.log.Error(err)
	}

	if err := gc.clearRepositories(); err != nil {
		gc.log.Error(err)
	}

	if full {
		gc.digestsClear()
	}
}

func (gc *GC) digestsClear() {
}
