package upload

import (
	"fmt"
	"github.com/MigAru/poseidon/internal/file_system"
)

func (m *Manager) UploadManifest(project, reference, hashType string, data []byte) (string, error) {
	var hash string
	hasher, err := m.hasher.Build(hashType, data)
	if err != nil {
		return hash, err
	}
	hash = fmt.Sprintf("%s:%x", hashType, hasher.Sum(nil))

	params := file_system.NewCreateParamsManifest(project, reference)
	if err := m.uploads.fs.CreateManifest(params.WithFilename(hash).WithData(data)); err != nil {
		return hash, err
	}

	if err := m.uploads.fs.CreateDigest(project, hash, data); err != nil {
		return hash, err
	}
	return hash, nil
}
