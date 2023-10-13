package migrate

import (
	"fmt"
	"github.com/caarlos0/env/v9"
)

type Config struct {
	driver string `env:"DRIVER" envDefault:""`
	dsn    string `env:"DSN" envDefault:""`
}

func newConfig() *Config {
	var cfg Config
	opts := env.Options{OnSet: func(tag string, value interface{}, isDefault bool) {
		fmt.Printf("Env value set %s to '%v' | is default - %v\n", tag, value, isDefault)
	}}

	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		return nil
	}

	return &cfg
}
