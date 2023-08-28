package upload

func (m *Manager) Update(id string, chunk []byte) error {
	params := NewUpdateParams(id).WithChunk(chunk).WithUploadedBytes(len(chunk))
	return m.uploads.update(params)
}
