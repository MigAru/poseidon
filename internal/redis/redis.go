package redis

import (
	"context"
	"errors"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/redis/simple"
)

type Redis interface {
	GetMarkedFiles(ctx context.Context, filetype string) ([]string, error)
	SetMarkFile(ctx context.Context, filetype string, name string) error
}

func New(cfg *config.Config) (Redis, func(), error) {
	switch cfg.Redis.Cluster {
	case true:
		panic("redis cluster not implemented")
	default:
		return simple.NewRedis(cfg.Redis.DSN)
	}

	return nil, func() {}, errors.New("get redis cluster switch")
}
