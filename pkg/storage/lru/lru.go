package lru

import (
	"errors"
	lruGo "github.com/hashicorp/golang-lru/v2"
	lruGoExpirable "github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/sirupsen/logrus"
	"time"
)

type LRU struct {
	log          *logrus.Logger
	WithTTL      bool
	cache        *lruGo.Cache[string, string]
	cacheWithTTL *lruGoExpirable.LRU[string, string]
}

func New(log *logrus.Logger, size int, ttl *time.Duration) (*LRU, error) {
	var (
		lru = LRU{log: log}
	)

	if size <= 0 {
		return nil, errors.New("suze must be not null")
	}

	if ttl != nil {
		cache := lruGoExpirable.NewLRU[string, string](size, nil, *ttl)
		lru.cacheWithTTL, lru.WithTTL = cache, true

	} else {
		cache, err := lruGo.New[string, string](size)
		if err != nil {
			return nil, err
		}
		lru.cache = cache
	}
	return &lru, nil
}

func (l *LRU) Get(key string) (string, error) {
	if l.WithTTL {
		value, ok := l.cacheWithTTL.Get(key)
		if !ok {
			//todo: вынести в отдельную ошибку
			return value, errors.New("element not found")
		}
		return value, nil
	}
	value, ok := l.cache.Get(key)
	if !ok {
		return value, errors.New("element not found")
	}

	return value, nil
}

func (l *LRU) Update(key, value string) error {

	if l.WithTTL {
		if ok := l.cacheWithTTL.Remove(key); !ok {
			return errors.New("element not found ")
		}
		ok := l.cacheWithTTL.Add(key, value)
		if !ok {
			return errors.New("")
		}
		return nil
	}
	if ok := l.cache.Remove(key); !ok {
		return errors.New("element not found ")
	}
	ok := l.cache.Add(key, value)
	if !ok {
		return errors.New("")
	}
	return nil
}

func (l *LRU) Create(key, value string) error {
	if l.WithTTL {
		l.cacheWithTTL.Add(key, value)
		return nil
	}
	l.cache.Add(key, value)
	return nil
}

func (l *LRU) Delete(key string) error {
	if l.WithTTL {
		ok := l.cacheWithTTL.Remove(key)
		if !ok {
			return errors.New("")
		}
		return nil
	}
	ok := l.cache.Remove(key)
	if !ok {
		return errors.New("")
	}
	return nil
}
