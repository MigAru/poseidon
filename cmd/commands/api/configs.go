package api

import (
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/consts"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

var configsSet = wire.NewSet(
	ProvideConfigFromCliContext,
)

func ProvideConfigFromCliContext(c *cli.Context) *config.Config {
	//TODO: переехать на env парсинг структуры
	return &config.Config{
		DebugMode: c.Bool(consts.DebugMode),
		Server: config.Server{
			Port:                     c.String(consts.ServerPort),
			TimeoutGracefullShutdown: c.Int(consts.ServerTimeout),
		},
	}
}
