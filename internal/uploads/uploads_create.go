package uploads

import (
	"github.com/google/uuid"
)

func (u *Uploads) Create() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return id.String(), err
	}

	return id.String(), u.fs.UploadBlob(id.String(), []byte{})
}
