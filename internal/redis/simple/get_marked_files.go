package simple

import (
	"context"
	"encoding/json"
)

func (r *Redis) GetMarkedFiles(ctx context.Context, filetype string) ([]string, error) {
	filesRaw, err := r.conn.Get(ctx, "marked_files_"+filetype).Result()
	if err != nil {
		return nil, err
	}
	var files []string
	if err := json.Unmarshal([]byte(filesRaw), &files); err != nil {
		return nil, err
	}

	return files, nil
}
