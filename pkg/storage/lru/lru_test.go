package lru

import (
	"github.com/MigAru/poseidon/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestImplementStorageInterface(t *testing.T) {
	assert.Implements(t, (*storage.ST)(nil), new(LRU), "interface realisation test")
}

func TestTTL(t *testing.T) {
	ttl := time.Millisecond * 1
	lru, err := New(logrus.New(), 1, &ttl)
	assert.NoError(t, err)

	err = lru.Create("test", "test")
	assert.NoError(t, err)

	timer := time.NewTimer(time.Millisecond * 5)
	<-timer.C

	//because not testing get
	keys := lru.cacheWithTTL.Keys()
	assert.EqualValues(t, 0, len(keys))

}

func TestCreateWithoutTTL(t *testing.T) {
	lru, err := New(logrus.New(), 1, nil)
	assert.NoError(t, err, "err create lru assert")

	err = lru.Create("test", "test")
	assert.NoError(t, err, "err create element assert")

	//because not testing get
	keys := lru.cache.Keys()
	assert.EqualValues(t, 1, len(keys), "")
}

func TestCreateWithTTL(t *testing.T) {
	ttl := time.Second * 2
	lru, err := New(logrus.New(), 1, &ttl)
	assert.NoError(t, err)

	err = lru.Create("test", "test")
	assert.NoError(t, err)

	//because not testing get
	keys := lru.cacheWithTTL.Keys()
	assert.EqualValues(t, 1, len(keys))
}

func TestGetWithoutTTL(t *testing.T) {
	var expected = "test"
	lru, err := New(logrus.New(), 1, nil)
	assert.NoError(t, err, "err create lru assert")

	err = lru.Create("test", expected)
	assert.NoError(t, err, "err create element assert")

	value, err := lru.Get("test")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, value)

}

func TestGetWithTTL(t *testing.T) {

}

func TestEvictedWithoutTTL(t *testing.T) {
	var expected = "evicted"

	lru, err := New(logrus.New(), 1, nil)
	assert.NoError(t, err, "err create lru assert")

	err = lru.Create("test", "test")
	assert.NoError(t, err, "err create element assert")

	err = lru.Create("evicted", expected)
	assert.NoError(t, err, "err create element assert")

	value, err := lru.Get("evicted")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, value)
}

func TestEvictedWithTTL(t *testing.T) {
	var expected = "evicted"
	ttl := time.Second * 2
	lru, err := New(logrus.New(), 1, &ttl)
	assert.NoError(t, err)

	err = lru.Create("test", "test")
	assert.NoError(t, err, "err create element assert")

	err = lru.Create("evicted", expected)
	assert.NoError(t, err, "err create element assert")

	value, err := lru.Get("evicted")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, value)
}

func TestDelete(t *testing.T) {}

func TestUpdate(t *testing.T) {}
