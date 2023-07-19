package api

import (
	"github.com/MigAru/poseidon/internal/interfaces/blob"
	"github.com/MigAru/poseidon/internal/interfaces/digest/digest"
	"github.com/MigAru/poseidon/internal/interfaces/manifest"
	repository2 "github.com/MigAru/poseidon/internal/registry/blob/repository"
	"github.com/MigAru/poseidon/internal/registry/digest/repository"
	repository3 "github.com/MigAru/poseidon/internal/registry/manifest/repository"
	"github.com/google/wire"
)

var dbSet = wire.NewSet(
	ProvideFileSystemBlobRepository,
	wire.Bind(new(blob.Repository), new(*repository2.FileSystem)),
	ProvideFileSystemDigestRepository,
	wire.Bind(new(digest.Repository), new(*repository.FileSystem)),
	ProvideFileSystemManifestRepository,
	wire.Bind(new(manifest.Repository), new(*repository3.FileSystem)),
)

func ProvideFileSystemBlobRepository() *repository2.FileSystem {
	//TODO: сделать переменную в конфиге где прописывается куда складывать файлы
	return repository2.NewFileSystem("tmp")
}

func ProvideFileSystemDigestRepository() *repository.FileSystem {
	return repository.NewFileSystem("tmp")
}

func ProvideFileSystemManifestRepository() *repository3.FileSystem {
	return repository3.NewFileSystem("tmp")
}
