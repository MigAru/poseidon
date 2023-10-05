package upload

import (
	"github.com/google/uuid"
	"time"
)

func (u *Uploads) Create(params CreateParams) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return id.String(), err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	uploadParams := NewInitUploadParams(id.String(), params.ProjectName).
		WithFS(u.fs).
		WithTotalSize(params.TotalSize).
		WithTimeout(params.Timeout)

	upload := InitUpload(uploadParams)
	u.unsafe[id.String()] = upload

	return id.String(), nil
}

type CreateParams struct {
	TotalSize   int
	ProjectName string
	Timeout     time.Duration
}
