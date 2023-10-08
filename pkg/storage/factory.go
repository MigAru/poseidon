package storage

import (
	"github.com/MigAru/poseidon/pkg/storage/lru"
	"github.com/sirupsen/logrus"
	"time"
)

type STFactory struct {
	log  *logrus.Logger
	size int
	ttl  *time.Duration
}

func NewSTFactory(log *logrus.Logger, size int, ttl *time.Duration) *STFactory {
	return &STFactory{
		log:  log,
		size: size,
		ttl:  ttl,
	}
}

func (stf *STFactory) Build(typeST int) *ST {
	switch typeST {
	case LRU:
		lru.New(stf.log, stf.size, stf.ttl)
	}
	panic("not supported type of ST")
}
