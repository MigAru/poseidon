package upload

func (m *Manager) Get(id string) (*Upload, bool) {
	return m.uploads.Get(id)
}
