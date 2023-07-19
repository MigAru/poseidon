package registry_api

import (
	"poseidon/internal/config"
	"poseidon/internal/consts"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

var configsSet = wire.NewSet(
	ProvideConfigFromCliContext,
)

func ProvideConfigFromCliContext(c *cli.Context) *config.Config {
	return &config.Config{
		DebugMode: c.Bool(consts.DebugMode),
		Server: config.Server{
			Port:                     c.String(consts.ServerPort),
			TimeoutGracefullShutdown: c.Int(consts.ServerTimeout),
		},
	}
}
