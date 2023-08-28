package upload

func (u *Upload) Done(digest string, finalChunk []byte) error {
	blobBytes, err := u.fs.GetBlob(u.ID)
	if err != nil {
		return err
	}
	blobBytes = append(blobBytes, finalChunk...)
	//so that the worker does not delay ahead of create digest
	if err := u.fs.CreateDigest(u.ProjectName, digest, blobBytes); err != nil {
		return err
	}
	u.CancelFunc()
	return nil
}
