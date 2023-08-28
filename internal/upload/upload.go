package upload

import (
	"github.com/MigAru/poseidon/internal/file_system"
)

type Upload struct {
	ID            string
	ProjectName   string
	fs            *file_system.FS
	TotalSize     int
	UploadedBytes int
	Bus           chan Chunk
}

type Chunk struct {
	UploadID string
	Bytes    []byte
}

type InitUploadParams struct {
	ID          string
	ProjectName string
	TotalSize   int
	Bus         chan Chunk
	FS          *file_system.FS
}

func NewInitUploadParams(id, projectName string, bus chan Chunk) *InitUploadParams {
	return &InitUploadParams{
		ID:          id,
		ProjectName: projectName,
		Bus:         bus,
	}
}

func (p *InitUploadParams) WithFS(fs *file_system.FS) *InitUploadParams {
	p.FS = fs
	return p
}

func NewUpload(params *InitUploadParams) *Upload {
	return &Upload{
		ID:          params.ID,
		fs:          params.FS,
		ProjectName: params.ProjectName,
		TotalSize:   params.TotalSize,
	}
}
