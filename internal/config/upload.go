package config

type Upload struct {
	BusWorkersLimit int `env:"BUS_WORKERS_LIMIT" envDefault:"1"`
}
