package simple

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"slices"
)

func (r *Redis) SetMarkFile(ctx context.Context, filetype string, name string) error {
	return r.conn.Watch(ctx, func(tx *redis.Tx) error {
		filesRaw, err := tx.Get(ctx, "marked_files_"+filetype).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}

		var files []string
		if !errors.Is(err, redis.Nil) {
			if err := json.Unmarshal([]byte(filesRaw), &files); err != nil {
				return err
			}
		}

		if !slices.Contains(files, name) {
			files = append(files, name)
		}

		setFilesRaw, err := json.Marshal(&files)
		if err != nil {
			return err
		}

		return tx.Set(ctx, "marked_files_"+filetype, string(setFilesRaw), 0).Err()
	})
}
