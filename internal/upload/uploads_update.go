package upload

import "errors"

func (u *Uploads) update(params *UpdateParams) error {
	u.mu.Lock()
	upload, ok := u.unsafe[params.ID]
	u.mu.Unlock()
	if !ok {
		return errors.New("upload not found")
	}

	if params.Chunk != nil {
		return u.fs.UploadBlob(upload.ID, params.Chunk)
	}
	if params.UploadedBytes > 0 {
		upload.UploadedBytes = upload.UploadedBytes + params.UploadedBytes
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
