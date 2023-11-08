package simple

import "github.com/redis/go-redis/v9"

type Redis struct {
	conn *redis.Client
}

func NewRedis(dsn string) (*Redis, func(), error) {
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, func() {}, err
	}
	client := redis.NewClient(opts)
	return &Redis{conn: client}, func() { client.Close() }, nil
}
