package locker

import (
	"encoding/json"
	"github.com/MigAru/poseidon/pkg/storage"
	"github.com/sirupsen/logrus"
	"time"
)

//LK - need for lock download image to delete image
type LK struct {
	log     logrus.Logger
	storage storage.ST
}

type Metadata struct {
	CreatedAt time.Time `json:"created_at"`
}

func (lk *LK) Lock(reference string) error {
	metadata := Metadata{CreatedAt: time.Now()}
	body, err := json.Marshal(&metadata)
	if err != nil {
		return err
	}
	return lk.storage.Create(reference, string(body))
}

func (lk *LK) Unlock(reference string) error {
	return lk.storage.Delete(reference)
}

func (lk *LK) Status(reference string) (*Metadata, error) {
	body, err := lk.storage.Get(reference)
	if err != nil {
		return nil, err
	}

	var metadata = new(Metadata)
	if err := json.Unmarshal([]byte(body), metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}
