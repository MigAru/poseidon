package upload

import "errors"

func (u *Uploads) Delete(id string) error {
	return u.delete(id)
}

func (u *Uploads) delete(id string) error {
	if _, ok := u.Get(id); !ok {
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
