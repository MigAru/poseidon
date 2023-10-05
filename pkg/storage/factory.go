package storage

import (
	"github.com/sirupsen/logrus"
	"time"
)

type STFactory struct {
	log  *logrus.Logger
	size int
	ttl  time.Duration
}

func NewSTFactory(log *logrus.Logger, size int, ttl time.Duration) *STFactory {
	return &STFactory{
		log:  log,
		size: size,
		ttl:  ttl,
	}
}

func (stf *STFactory) Build(typeST int) *ST {
	switch typeST {
	case LRU:

	case Redis:
	}
	panic("not supported type of ST")
}
