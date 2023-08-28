package upload

import (
	"github.com/MigAru/poseidon/internal/file_system"
	"sync"
)

type Uploads struct {
	mu     sync.RWMutex
	fs     *file_system.FS
	bus    chan Chunk
	unsafe map[string]*Upload
}

func NewUploads() *Uploads {
	return &Uploads{
		unsafe: make(map[string]*Upload),
		bus:    make(chan Chunk),
	}
}
