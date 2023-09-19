package upload

func (m *Manager) DeleteUpload(id string) error {
	return m.uploads.delete(id)
}
