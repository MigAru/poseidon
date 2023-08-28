package upload

import "errors"

func (u *Uploads) Update(params *UpdateParams) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	upload, ok := u.unsafe[params.ID]
	if !ok {
		return errors.New("upload not found")
	}

	if params.Chunk != nil {
		upload.Bus <- params.Chunk
	}
	if params.UploadedBytes > 0 {
		upload.UploadedBytes += params.UploadedBytes
	}
	return nil
}

type UpdateParams struct {
	ID            string
	UploadedBytes int
	Chunk         []byte
}

func NewUpdateParams(id string) *UpdateParams {
	return &UpdateParams{ID: id}
}

func (p *UpdateParams) WithUploadedBytes(size int) *UpdateParams {
	p.UploadedBytes = size
	return p
}

func (p *UpdateParams) WithChunk(chunk []byte) *UpdateParams {
	p.Chunk = chunk
	return p
}
