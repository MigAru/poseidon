package upload

import "sync"

type Uploads struct {
	mu     sync.RWMutex
	unsafe map[string]*Upload
}

func NewUploads() *Uploads {
	return &Uploads{
		unsafe: make(map[string]*Upload),
	}
}
