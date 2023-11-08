package upload

func (u *Uploads) Delete(id string) error {
	if _, err := u.fs.GetBlob(id); err != nil {
		return err
	}

	if err := u.fs.DeleteBlob(id); err != nil {
		return err
	}

	return nil
}
