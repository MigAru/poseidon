package blob

import (
	"bytes"
	"io"
	"os"
	"path"
)

type FileSystemRepository struct {
	basePath string
}

func NewFileSystemRepository(basePath string) *FileSystemRepository {
	basePath = path.Join(basePath, "blobs")
	return &FileSystemRepository{basePath: basePath}
}
func (r FileSystemRepository) normalizePath(name string) string {
	return path.Join(r.basePath, name)
}

func (r FileSystemRepository) Get(name string) ([]byte, error) {
	return os.ReadFile(r.normalizePath(name))
}

func (r FileSystemRepository) Create(name string, data []byte) error {
	name = r.normalizePath(name)
	err := os.MkdirAll(r.basePath, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	perm := os.O_CREATE
	if fileExist(name) {
		perm = perm | os.O_APPEND | os.O_WRONLY
	} else {
		perm = perm | os.O_RDWR
	}
	f, err := os.OpenFile(name, perm, 0750)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewBuffer(data))
	if err != nil {
		os.Remove(name)
		return err
	}
	return nil
}

func fileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (r FileSystemRepository) Exist(name string) bool {
	if _, err := os.Stat(r.normalizePath(name)); err != nil {
		return false
	}
	return true
}

func (r FileSystemRepository) Delete(name string) error {
	return os.Remove(r.normalizePath(name))
}
