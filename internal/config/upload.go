package config

type Upload struct {
	ChunkSize      int `env:"CHUNK_SIZE" envDefault:"500"` //in bytes
	MinutesTimeout int `env:"MINUTES_TIMEOUT" envDefault:"5"`
}
