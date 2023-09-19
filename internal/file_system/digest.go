package file_system

import (
	"bytes"
	"io"
	"os"
	"path"
	"strings"
)

func (f *FS) GetDigest(project, name string) ([]byte, error) {
	ar := strings.Split(name, ":")
	algo, hash := ar[0], ar[1]
	digestPath := path.Join(f.basePath, "repos", project, algo, hash[:3], name)
	return os.ReadFile(digestPath)
}

func (f *FS) CreateDigest(project, name string, data []byte) error {
	ar := strings.Split(name, ":")
	algo, hash := ar[0], ar[1]
	digestPath := path.Join(f.basePath, "repos", project, algo, hash[:3])
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
	file, err := os.OpenFile(path.Join(digestPath, name), perm, 0750)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewBuffer(data))
	if err != nil {
		if err := os.Remove(name); err != nil {
			return err
		}
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

func (f *FS) ExistDigest(project, name string) error {
	//TODO implement me
	panic("implement me")
}

func (f *FS) DeleteDigest(project, name string) error {
	//TODO implement me
	panic("implement me")
}
