package upload

import "github.com/google/uuid"

func (u *Uploads) Create(params CreateParams) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return id.String(), err
	}
	u.mu.Lock()
	defer u.mu.Unlock()

	u.unsafe[id.String()] = NewUpload(id.String(), params.ProjectName, params.TotalSize)

	return id.String(), nil
}

type CreateParams struct {
	TotalSize   int
	ProjectName string
}
