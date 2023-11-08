package upload

func (u *Uploads) Done(id string, digest string, chunk []byte) (int, error) {
	blobBytes, err := u.fs.GetBlob(id)
	if err != nil {
		return 0, err
	}

	blobBytes = append(blobBytes, chunk...)

	if err := u.fs.CreateDigest(digest, blobBytes); err != nil {
		return 0, err
	}

	return len(blobBytes), nil
}
