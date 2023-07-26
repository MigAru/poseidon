package hasher

import (
	"crypto/sha256"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/MigAru/poseidon/pkg/registry/hasher/methods"
	"hash"
)

type Factory struct {
	method string
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Build() (hash.Hash, error) {
	switch f.method {
	case methods.SHA256:
		return sha256.New(), nil
	}
	return nil, errors.NotValidMethod
}
