package manifest

import "github.com/MigAru/poseidon/internal/file_system"

type Manager struct {
	fs *file_system.FS
}

func NewManager(fs *file_system.FS) *Manager {
	return &Manager{fs: fs}
}
