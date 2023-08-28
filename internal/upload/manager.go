package upload

import (
	"context"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/sirupsen/logrus"
)

type Manager struct {
	ctx             context.Context
	log             *logrus.Entry
	uploads         *Uploads
	bus             chan Metadata
	chunkBus        chan Chunk
	busWorkersLimit int
}

type Metadata struct {
	ID            string
	UploadedBytes int
}

func NewManager(ctx context.Context, cfg *config.Config, log *logrus.Logger) *Manager {
	return &Manager{
		ctx:             ctx,
		log:             log.WithField("prefix", "upload_manager"),
		uploads:         NewUploads(),
		bus:             make(chan Metadata),
		busWorkersLimit: cfg.Upload.BusWorkersLimit,
	}
}

func (m *Manager) StartBusWorkers() {
	log := m.log.WithField("second_prefix", "worker_bus")
	for i := 1; i >= m.busWorkersLimit; i++ {
		go func() {

			for {
				select {
				case <-m.ctx.Done():
					break
				case metadata := <-m.bus:
					log.Infof("getting metadata from bus | upload id: %s", metadata.ID)
					params := NewUpdateParams(metadata.ID).WithUploadedBytes(metadata.UploadedBytes)
					if err := m.uploads.Update(params); err != nil {
						log.Errorf("status update: %s | upload id: %s", err.Error(), metadata.ID)
						continue
					}
					log.Infof("status update: %s | upload id: %s", "success", metadata.ID)
				}
			}
		}()
	}
}
