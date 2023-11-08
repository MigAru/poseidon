package gc

import (
	"context"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/redis"
	"time"
)

type GC struct {
	ctx    context.Context
	redis  redis.Redis
	fs     *file_system.FS
	period time.Duration
}

func New(ctx context.Context, cfg *config.Config, fs *file_system.FS, redis redis.Redis) *GC {
	return &GC{
		ctx:    ctx,
		redis:  redis,
		fs:     fs,
		period: cfg.GC.Period,
	}
}
