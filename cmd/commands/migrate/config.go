package migrate

import (
	"fmt"
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Driver string `env:"DATABASE_DRIVER" envDefault:"sqlite3"`
	DSN    string `env:"DATABASE_DSN" envDefault:"test.db"`
}

func newConfig() *Config {
	var cfg = &Config{}
	opts := env.Options{OnSet: func(tag string, value interface{}, isDefault bool) {
		fmt.Printf("Env value set %s to '%v' | is default - %v\n", tag, value, isDefault)
	}}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil
	}

	return cfg
}
