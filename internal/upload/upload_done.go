package upload

import (
	"errors"
)

func (u *Upload) Done(digest string, finalChunk []byte) (int, error) {
	blobBytes, err := u.fs.GetBlob(u.ID)
	if err != nil {
		return 0, err
	}

	blobBytes = append(blobBytes, finalChunk...)
	if u.TotalSize != len(blobBytes) {
		return 0, errors.New("short download")
	}

	if err := u.fs.CreateDigest(u.ProjectName, digest, blobBytes); err != nil {
		return 0, err
	}

	defer u.cancel()

	return len(blobBytes), nil
}
