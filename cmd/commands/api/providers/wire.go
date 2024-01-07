//go:build wireinject
// +build wireinject

package providers

import (
	"context"
	"github.com/MigAru/poseidon/internal/base"
	"github.com/MigAru/poseidon/internal/blob"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/database"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/gin"
	"github.com/MigAru/poseidon/internal/logger"
	"github.com/MigAru/poseidon/internal/manifest"
	"github.com/MigAru/poseidon/internal/ping"
	"github.com/MigAru/poseidon/internal/tech"
	"github.com/MigAru/poseidon/internal/uploads"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/hasher"
	"github.com/google/wire"
)

type Backend struct{}

func InitializeBackend(ctx context.Context) (Backend, func(), error) {
	panic(wire.Build(
		config.NewFromEnv,
		logger.NewLogrus,
		database.New,
		file_system.New,
		hasher.New,
		uploads.NewUploads,
		ping.NewController,
		base.NewController,
		blob.NewController,
		manifest.NewController,
		tech.NewController,
		gin.NewServer,
		wire.Bind(new(http.Server), new(*gin.Server)),
		ServiceProvider,
	))
}
