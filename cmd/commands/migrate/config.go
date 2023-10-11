package migrate

type Config struct {
	driver string `env:"DRIVER"`
	dsn    string `env:"DSN"`
}
