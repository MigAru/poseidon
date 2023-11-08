package file_system

import (
	"bytes"
	"io"
	"os"
	"path"
	"strings"
)

func (f *FS) GetDigest(digest string) ([]byte, error) {
	ar := strings.Split(digest, ":")
	algo, hash := ar[0], ar[1]
	digestPath := path.Join(f.basePath, "digest", algo, hash[:3], digest)
	return os.ReadFile(digestPath)
}

func (f *FS) CreateDigest(digest string, data []byte) error {
	ar := strings.Split(digest, ":")
	algo, hash := ar[0], ar[1]
	digestPath := path.Join(f.basePath, "digest", algo, hash[:3])
	err := os.MkdirAll(digestPath, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	perm := os.O_CREATE
	if fileExist(path.Join(digestPath, digest)) {
		perm = perm | os.O_TRUNC | os.O_WRONLY
	} else {
		perm = perm | os.O_RDWR
	}
	file, err := os.OpenFile(path.Join(digestPath, digest), perm, 0750)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewBuffer(data))
	if err != nil {
		if err := os.Remove(digest); err != nil {
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

func (f *FS) DeleteDigest(digest string) error {
	ar := strings.Split(digest, ":")
	algo, hash := ar[0], ar[1]
	digestPath := path.Join(f.basePath, "digest", algo, hash[:3], digest)
	return os.Remove(digestPath)
}
