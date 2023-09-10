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

func New() *Hasher {
	return &Hasher{}
}

func (f *Hasher) Build(method string, data []byte) (hash.Hash, error) {
	switch method {
	case methods.SHA224:
		return f.createHasher(sha256.New224(), data)
	case methods.SHA256:
		return f.createHasher(sha256.New(), data)
	case methods.SHA384:
		return f.createHasher(sha512.New384(), data)
	case methods.SHA512:
		return f.createHasher(sha512.New(), data)
	case methods.MD5:
		return f.createHasher(md5.New(), data)
	}
	return nil, errors.NotSupportedMethod
}

func (f *Hasher) createHasher(hasher hash.Hash, data []byte) (hash.Hash, error) {
	if _, err := hasher.Write(data); err != nil {
		return nil, err
	}
	return hasher, nil
}
