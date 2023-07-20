package config

type Config struct {
	DebugMode bool   `env:"DEBUG_MODE" envDefault:"true"`
	Server    Server `envPrefix:"SERVER_"`
}
