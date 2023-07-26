package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/MigAru/poseidon/pkg/registry/hasher/methods"
	"hash"
)

type Hasher struct {
}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (f *Hasher) Build(method string) (hash.Hash, error) {
	switch method {
	case methods.SHA224:
		return sha256.New224(), nil
	case methods.SHA256:
		return sha256.New(), nil
	case methods.SHA384:
		return sha512.New384(), nil
	case methods.SHA512:
		return sha512.New(), nil
	case methods.MD5:
		return md5.New(), nil
	}
	return nil, errors.NotSupportedMethod
}
