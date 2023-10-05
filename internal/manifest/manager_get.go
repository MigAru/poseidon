package manifest

import (
	"encoding/json"
	"github.com/MigAru/poseidon/internal/file_system"
	v2_2 "github.com/MigAru/poseidon/pkg/registry/manifest/schema/v2.2"
	"strings"
)

func (m *Manager) Get(project, reference string) (v2_2.Manifest, string, error) {
	var (
		manifest  v2_2.Manifest
		filename  = reference
		fileBytes []byte
		err       error
	)
	params := file_system.NewGetParamsManifest(project, reference)
	if !m.isDigest(reference) {
		filename, err = m.fs.GetManifest(params)
		if err != nil {
			return manifest, filename, err
		}
	}
	fileBytes, err = m.fs.GetDigest(project, filename)
	if err != nil {
		return manifest, filename, err
	}
	//TODO: сделать универсальный unmarshaler для manifest v2 v1/oci/manifest list v2
	if err := json.Unmarshal(fileBytes, &manifest); err != nil {
		return manifest, filename, err
	}
	return manifest, filename, nil
}
func (m *Manager) isDigest(name string) bool {
	hashArray := strings.Split(name, ":")
	return len(hashArray) > 1
}