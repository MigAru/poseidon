package file_system

import (
	"io/fs"
	"path/filepath"
	"time"
)

func (f *FS) GetBlobsForDelete() ([]string, error) {
	var fileNames []string
	if err := filepath.Walk(f.normalizePath("/"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.ModTime().Add(f.blobExpire).Before(time.Now()) && !info.IsDir() {
			fileNames = append(fileNames, info.Name())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return fileNames, nil
}
