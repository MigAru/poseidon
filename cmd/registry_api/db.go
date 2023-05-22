package registry_api

import (
	"github.com/google/wire"
	blob2 "poseidon/internal/registry/blob"
	digest2 "poseidon/internal/registry/digest"
	repository2 "poseidon/pkg/registry/blob/repository"
	"poseidon/pkg/registry/digest/repository"
)

var dbSet = wire.NewSet(
	ProvideFileSystemBlobRepository,
	wire.Bind(new(blob2.Repository), new(*repository2.FileSystem)),
	ProvideFileSystemDigestRepository,
	wire.Bind(new(digest2.Repository), new(*repository.FileSystem)),
)

func ProvideFileSystemBlobRepository() *repository2.FileSystem {
	return repository2.NewFileSystem("tmp")
}

func ProvideFileSystemDigestRepository() *repository.FileSystem {
	return repository.NewFileSystem("tmp")
}
