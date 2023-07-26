package config

type FileSystem struct {
	BasePath string `env:"BASE_PATH" envDefault:"./tmp"`
}
