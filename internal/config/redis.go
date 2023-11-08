package config

type Redis struct {
	Cluster bool   `env:"CLUSTER" envDefault:"false"`
	DSN     string `env:"DSN" envDefault:"redis://localhost:6379"`
}
