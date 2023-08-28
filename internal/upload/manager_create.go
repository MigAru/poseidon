package upload

import "context"

func (m *Manager) Create(ctx context.Context, projectName string, TotalSize int) (string, error) {
	params := CreateParams{ProjectName: projectName, TotalSize: TotalSize, Timeout: m.defaultTimeout}

	return m.uploads.Create(ctx, params)
}
