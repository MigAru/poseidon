package upload

func (m *Manager) Delete(id string) error {
	return m.uploads.delete(id)
}
