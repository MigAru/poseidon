package providers

import (
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/interfaces/digest/digest"
	"github.com/MigAru/poseidon/internal/interfaces/manifest"
	"github.com/MigAru/poseidon/internal/registry/digest/repository"
	repository3 "github.com/MigAru/poseidon/internal/registry/manifest/repository"
	"github.com/google/wire"
)

var dbSet = wire.NewSet(
	ProvideFileSystem,
	ProvideFileSystemDigestRepository,
	wire.Bind(new(digest.Repository), new(*repository.FileSystem)),
	ProvideFileSystemManifestRepository,
	wire.Bind(new(manifest.Repository), new(*repository3.FileSystem)),
)

func ProvideFileSystem(_ *config.Config) *file_system.FS {
	return file_system.New("tmp")
}

func ProvideFileSystemDigestRepository() *repository.FileSystem {
	return repository.NewFileSystem("tmp")
}

func ProvideFileSystemManifestRepository() *repository3.FileSystem {
	return repository3.NewFileSystem("tmp")
}
