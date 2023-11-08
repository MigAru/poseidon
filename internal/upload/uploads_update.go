package upload

func (u *Uploads) Update(params *UpdateParams) error {
	if _, err := u.fs.GetBlob(params.ID); err != nil {
		return err
	}

	if params.Chunk != nil {
		return u.fs.UploadBlob(params.ID, params.Chunk)
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

func (p *UpdateParams) WithChunk(chunk []byte) *UpdateParams {
	p.Chunk = chunk
	return p
}
