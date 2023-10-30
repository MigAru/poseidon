package upload

import (
	"fmt"
)

func (m *Manager) UploadManifest(project, hashType string, data []byte) (string, error) {
	var hash string
	hasher, err := m.hasher.Build(hashType, data)
	if err != nil {
		return hash, err
	}
	hash = fmt.Sprintf("%s:%x", hashType, hasher.Sum(nil))

	if err := m.uploads.fs.CreateDigest(project, hash, data); err != nil {
		return hash, err
	}
	return hash, nil
}
