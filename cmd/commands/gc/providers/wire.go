//go:build wireinject
// +build wireinject

package providers

import (
	"context"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/gc"
	"github.com/MigAru/poseidon/internal/redis"
	"github.com/google/wire"
)

type App struct{}

func InitializeApp(ctx context.Context) (App, func(), error) {
	panic(wire.Build(
		config.NewFromEnv,
		redis.New,
		file_system.New,
		gc.New,
		AppProvider,
	))
}
