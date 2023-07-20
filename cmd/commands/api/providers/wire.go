//go:build wireinject
// +build wireinject

package providers

import (
	"context"
	"github.com/google/wire"
)

type Backend struct{}

func InitializeBackend(ctx context.Context) (Backend, func(), error) {
	panic(wire.Build(
		configsSet,
		loggersSet,
		dbSet,
		httpSet,
		servicesSet,
	))
}
