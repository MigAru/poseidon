package upload

func (u *Uploads) Get(id string) ([]byte, error) {
	return u.fs.GetBlob(id)
}
