package repository

import (
	"bytes"
	"io"
	"os"
	"path"
)

type FileSystem struct {
	basePath string
}

func NewFileSystem(basePath string) *FileSystem {
	//TODO: сделать воркера чтобы удалял после определенного времени блобы(временные файлы)
	basePath = path.Join(basePath, "blobs")
	return &FileSystem{basePath: basePath}
}

func (r FileSystem) Get(name string) ([]byte, error) {
	return os.ReadFile(r.normalizePath(name))
}

func (r FileSystem) normalizePath(name string) string {
	return path.Join(r.basePath, name)
}

func (r FileSystem) Create(name string, data []byte) error {
	name = r.normalizePath(name)
	err := os.MkdirAll(r.basePath, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	perm := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
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

func (r FileSystem) Exist(name string) bool {
	if _, err := os.Stat(r.normalizePath(name)); err != nil {
		return false
	}
	return true
}

func (r FileSystem) Delete(name string) error {
	return os.Remove(r.normalizePath(name))
}
