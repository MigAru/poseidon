package config

type Database struct {
	Driver string `env:"DRIVER" envDefault:"sqlite3"`
	DSN    string `env:"DSN" envDefault:"test.db"`
}
