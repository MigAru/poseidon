package upload

func (u *Uploads) Delete(id string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.unsafe, id)
}
