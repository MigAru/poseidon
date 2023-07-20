package providers

import (
	"fmt"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/caarlos0/env/v9"
	"github.com/google/wire"
)

var configsSet = wire.NewSet(
	ProvideConfigFromEnv,
)

func ProvideConfigFromEnv() (*config.Config, error) {
	cfg := &config.Config{}
	opts := env.Options{OnSet: func(tag string, value interface{}, isDefault bool) {
		fmt.Printf("Env value set %s to '%v' | is default - %v\n", tag, value, isDefault)
	}}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, err
	}
	return cfg, nil
}
