package upload

import "errors"

func (u *Uploads) delete(id string) error {
	_, ok := u.Get(id)
	if !ok {
		return errors.New("upload has been deleted")
	}

	if err := u.fs.DeleteBlob(id); err != nil {
		return err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	delete(u.unsafe, id)
	return nil
}
