package upload

import "errors"

func (u *Uploads) delete(id string) error {
	upload, ok := u.Get(id)
	if !ok {
		return errors.New("upload has been deleted")
	}

	if err := u.fs.DeleteBlob(id); err != nil {
		return err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	defer upload.cancel()

	delete(u.unsafe, id)
	return nil
}
