package repository

import (
	"bytes"
	"io"
	"os"
	"path"
	"strings"
)

type FileSystem struct {
	basePath string
}

func NewFileSystem(basePath string) *FileSystem {
	basePath = path.Join(basePath, "digests")
	return &FileSystem{basePath: basePath}
}

func (r FileSystem) Get(project, name string) ([]byte, error) {
	ar := strings.Split(name, ":")
	algo, hash := ar[0], ar[1]
	digestPath := path.Join(r.basePath, project, algo, hash[:3], name)
	return os.ReadFile(digestPath)
}

func (r FileSystem) Create(project, name string, data []byte) error {
	ar := strings.Split(name, ":")
	algo, hash := ar[0], ar[1]
	digestPath := path.Join(r.basePath, project, algo, hash[:3])
	err := os.MkdirAll(digestPath, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	perm := os.O_CREATE
	if fileExist(path.Join(digestPath, name)) {
		perm = perm | os.O_APPEND | os.O_WRONLY
	} else {
		perm = perm | os.O_RDWR
	}
	f, err := os.OpenFile(path.Join(digestPath, name), perm, 0750)
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

func (r FileSystem) Exist(project, name string) error {
	//TODO implement me
	panic("implement me")
}

func (r FileSystem) Delete(project, name string) error {
	//TODO implement me
	panic("implement me")
}
