package manifest

import "github.com/MigAru/poseidon/internal/file_system"

func (m *Manager) Delete(project, reference string) error {
	var (
		err error
	)
	if !m.isDigest(reference) {
		params := file_system.NewGetParamsManifest(project, reference)
		reference, err = m.fs.GetManifest(params)
		if err != nil {
			return err
		}
	}

	err = m.fs.DeleteManifest(file_system.NewBaseParamsManifest(project, reference))
	if err != nil {
		return err
	}

	err = m.fs.DeleteDigest(project, reference)
	if err != nil {
		return err
	}
	return err
}
