package upload

func (m *Manager) Create(projectName string, TotalSize int) (string, error) {
	params := CreateParams{ProjectName: projectName, TotalSize: TotalSize}

	return m.uploads.Create(params)
}
