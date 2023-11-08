package config

import "time"

type GC struct {
	Period time.Duration `env:"PERIOD" envDefault:"1s"`
}
