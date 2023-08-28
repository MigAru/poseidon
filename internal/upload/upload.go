package upload

import (
	"context"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/sirupsen/logrus"
	"time"
)

type Upload struct {
	ID            string
	ProjectName   string
	Monotonic     bool
	ChunkNum      int //for calculate +- size get
	log           *logrus.Entry
	timeout       time.Duration //deadline routine
	ctx           context.Context
	cancel        context.CancelFunc
	fs            *file_system.FS
	TotalSize     int
	UploadedBytes int //for final check
	Errors        []error
	Queue         chan []byte
}

type Chunk struct {
	Bytes []byte
}

type InitUploadParams struct {
	ID          string
	Timeout     time.Duration
	Log         *logrus.Logger
	ProjectName string
	TotalSize   int
	FS          *file_system.FS
}

func NewInitUploadParams(id, projectName string) *InitUploadParams {
	return &InitUploadParams{
		ID:          id,
		ProjectName: projectName,
	}
}

func (p *InitUploadParams) WithTimeout(timeout time.Duration) *InitUploadParams {
	p.Timeout = timeout
	return p
}

func (p *InitUploadParams) WithFS(fs *file_system.FS) *InitUploadParams {
	p.FS = fs
	return p
}

func (p *InitUploadParams) WithTotalSize(size int) *InitUploadParams {
	p.TotalSize = size
	return p
}

func (p *InitUploadParams) WithLogger(log *logrus.Logger) *InitUploadParams {
	p.Log = log
	return p
}

func InitUpload(ctx context.Context, params *InitUploadParams) *Upload {
	uploadCTX, cancel := context.WithCancel(ctx)

	return &Upload{
		ID:          params.ID,
		timeout:     params.Timeout,
		log:         logrus.WithField("prefix", "upload"),
		fs:          params.FS,
		ctx:         uploadCTX,
		cancel:      cancel,
		ProjectName: params.ProjectName,
		TotalSize:   params.TotalSize,
		Queue:       make(chan []byte),
	}
}
