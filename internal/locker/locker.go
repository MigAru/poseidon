package locker

import (
	"github.com/MigAru/poseidon/pkg/storage"
	"github.com/sirupsen/logrus"
)

//LK - need for lock download image to delete image
type LK struct {
	log     logrus.Logger
	storage storage.ST
}

type Metadata struct{}

func (lk *LK) Lock(reference string) {}

func (lk *LK) Unlock(reference string) {}

func (lk *LK) Status(reference string) {}
