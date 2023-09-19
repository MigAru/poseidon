//go:build wireinject
// +build wireinject

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
	"github.com/google/wire"
)

type Backend struct{}

func InitializeBackend(ctx context.Context) (Backend, func(), error) {
	panic(wire.Build(
		config.NewFromEnv,
		logger.NewLogrus,
		file_system.New,
		hasher.New,
		upload.NewManager,
		ping.NewPingController,
		base.NewController,
		blob.NewController,
		manifest.NewManager,
		manifest.NewController,
		ServerProvider,
		ServiceProvider,
	))
}
