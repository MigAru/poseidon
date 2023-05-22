package registry_api

import (
	"github.com/google/wire"
	blob2 "poseidon/internal/registry/blob"
	digest2 "poseidon/internal/registry/digest"
	"poseidon/pkg/registry/blob"
	"poseidon/pkg/registry/digest"
)

var dbSet = wire.NewSet(
	ProvideFileSystemBlobRepository,
	wire.Bind(new(blob2.Repository), new(*blob.FileSystemRepository)),
	ProvideFileSystemDigestRepository,
	wire.Bind(new(digest2.Repository), new(*digest.FileSystemRepository)),
)

func ProvideFileSystemBlobRepository() *blob.FileSystemRepository {
	return blob.NewFileSystemRepository("tmp")
}

func ProvideFileSystemDigestRepository() *digest.FileSystemRepository {
	return digest.NewFileSystemRepository("tmp")
}
