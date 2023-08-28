package upload

import "github.com/google/uuid"

func (u *Uploads) Create(params CreateParams) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return id.String(), err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	uploadParams := NewInitUploadParams(id.String(), params.ProjectName, u.bus)
	upload := NewUpload(uploadParams.WithFS(u.fs))

	u.unsafe[id.String()] = upload

	return id.String(), nil
}

type CreateParams struct {
	TotalSize   int
	ProjectName string
	Bus         chan Chunk
}
