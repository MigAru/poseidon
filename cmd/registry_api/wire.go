//go:build wireinject
// +build wireinject

package registry_api

import (
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Backend struct{}

func InitializeBackend(ctx *cli.Context) (Backend, func(), error) {
	panic(wire.Build(
		configsSet,
		loggersSet,
		dbSet,
		httpSet,
		servicesSet,
	))
}
