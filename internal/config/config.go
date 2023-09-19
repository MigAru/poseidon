package config

import (
	"fmt"
	"github.com/caarlos0/env/v9"
)

type Config struct {
	DebugMode  bool       `env:"DEBUG_MODE" envDefault:"true"`
	FileSystem FileSystem `env:"FILE_SYSTEM_"`
	Server     Server     `envPrefix:"SERVER_"`
	Upload     Upload     `envPrefix:"UPLOAD_"`
}

func NewFromEnv() (*Config, error) {
	cfg := &Config{}
	opts := env.Options{OnSet: func(tag string, value interface{}, isDefault bool) {
		fmt.Printf("Env value set %s to '%v' | is default - %v\n", tag, value, isDefault)
	}}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, err
	}
	return cfg, nil
}
