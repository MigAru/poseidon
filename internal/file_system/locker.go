package file_system

import (
	"encoding/json"
	"github.com/MigAru/poseidon/pkg/storage"
	"github.com/sirupsen/logrus"
	"time"
)

//Locker - need for lock download image to delete image
type Locker struct {
	log     logrus.Logger
	storage storage.ST
}

func NewLocker(log logrus.Logger, st storage.ST) Locker {
	return Locker{log: log, storage: st}
}

//Metadata - for create queue or detect deadlock and more
type Metadata struct {
	CreatedAt time.Time `json:"created_at"` //time for detect deadlock(if ttl over big)
}

func (lk *Locker) Lock(reference string) error {
	metadata := Metadata{CreatedAt: time.Now()}
	body, err := json.Marshal(&metadata)
	if err != nil {
		return err
	}
	return lk.storage.Create(reference, string(body))
}

func (lk *Locker) Unlock(reference string) error {
	return lk.storage.Delete(reference)
}

func (lk *Locker) Status(reference string) (*Metadata, error) {
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
