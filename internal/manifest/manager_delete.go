package manifest

//TODO: убрать функцию после разработки garbage collector
func (m *Manager) Delete(project, reference string) error {
	var (
		err error
	)

	err = m.fs.DeleteDigest(project, reference)
	if err != nil {
		return err
	}
	return err
}
