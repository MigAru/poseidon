// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"github.com/urfave/cli/v2"
	"poseidon/internal/ping"
	"poseidon/internal/registry/base"
	"poseidon/internal/registry/blob"
	"poseidon/internal/registry/manifest"
)

// Injectors from wire.go:

func InitializeBackend(ctx *cli.Context) (Backend, func(), error) {
	config := ProvideConfigFromCliContext(ctx)
	logger, cleanup, err := ProvideNewLogger(config)
	if err != nil {
		return Backend{}, nil, err
	}
	pingController := ping.NewPingController()
	fileSystem := ProvideFileSystemBlobRepository()
	repositoryFileSystem := ProvideFileSystemDigestRepository()
	controller := blob.NewController(logger, fileSystem, repositoryFileSystem)
	baseController := base.NewController(logger)
	fileSystem2 := ProvideFileSystemManifestRepository()
	manifestController := manifest.NewController(logger, fileSystem2, repositoryFileSystem)
	httpServer := ServerProvider(config, logger, pingController, controller, baseController, manifestController)
	backend, err := BackendServiceProvider(httpServer)
	if err != nil {
		cleanup()
		return Backend{}, nil, err
	}
	return backend, func() {
		cleanup()
	}, nil
}

// wire.go:

type Backend struct{}
