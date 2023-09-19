package upload

func (u *Uploads) Get(id string) (upload *Upload, ok bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	upload, ok = u.unsafe[id]
	return
}
