package config

type Server struct {
	Port                    string `env:"PORT" envDefault:":8000"`
	TimeoutGracefulShutdown int    `env:"GRACEFUL_SHUTDOWN_TIMEOUT" envDefault:"15"`
}
