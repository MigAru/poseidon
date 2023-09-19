package file_system

import (
	"bytes"
	"io"
	"os"
	"path"
)

func (f *FS) GetBlob(name string) ([]byte, error) {
	return os.ReadFile(f.normalizePath(name))
}

func (f *FS) normalizePath(name string) string {
	return path.Join(f.basePath, "blobs", name)
}

func (f *FS) UploadBlob(name string, data []byte) error {
	name = f.normalizePath(name)
	err := os.MkdirAll(f.basePath+"/blobs", 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	perm := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(name, perm, 0750)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewBuffer(data))
	if err != nil {
		//dump error handle
		if err := os.Remove(name); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (f *FS) CheckExistBlob(name string) bool {
	if _, err := os.Stat(f.normalizePath(name)); err != nil {
		return false
	}
	return true
}

func (f *FS) DeleteBlob(_ string) error {
	//TODO: реализовать после менеджера загрузок
	panic("implement me")
}
