// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package providers

import (
	"context"
	"github.com/MigAru/poseidon/internal/base"
	"github.com/MigAru/poseidon/internal/blob"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/logger"
	"github.com/MigAru/poseidon/internal/manifest"
	"github.com/MigAru/poseidon/internal/ping"
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/MigAru/poseidon/pkg/registry/hasher"
)

// Injectors from wire.go:

func InitializeBackend(ctx context.Context) (Backend, func(), error) {
	configConfig, err := config.NewFromEnv()
	if err != nil {
		return Backend{}, nil, err
	}
	logrusLogger, cleanup, err := logger.NewLogrus(configConfig)
	if err != nil {
		return Backend{}, nil, err
	}
	pingController := ping.NewPingController()
	fs := file_system.New(configConfig)
	hasherHasher := hasher.New()
	manager := upload.NewManager(ctx, configConfig, fs, hasherHasher, logrusLogger)
	controller := blob.NewController(logrusLogger, configConfig, fs, manager)
	baseController := base.NewController(logrusLogger)
	manifestManager := manifest.NewManager(fs)
	manifestController := manifest.NewController(logrusLogger, fs, manager, manifestManager)
	server := ServerProvider(configConfig, logrusLogger, pingController, controller, baseController, manifestController)
	backend, cleanup2, err := ServiceProvider(ctx, server)
	if err != nil {
		cleanup()
		return Backend{}, nil, err
	}
	return backend, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type Backend struct{}
