package repository

import (
	"bytes"
	"errors"
	"github.com/MigAru/poseidon/internal/interfaces/manifest"
	"io"
	"os"
	"path"
)

type FileSystem struct {
	basePath string
}

func NewFileSystem(basePath string) *FileSystem {
	basePath = path.Join(basePath, "manifests")
	return &FileSystem{basePath: basePath}
}

func (r FileSystem) Get(params *manifest.GetParams) (filename string, err error) {
	manifestsPath := path.Join(r.basePath, params.Project, params.Tag)
	if params.Filename == "" {
		filename, err = r.getFilenameFromDir(manifestsPath)
		if err != nil {
			return filename, err
		}
	}
	bytesFile, err := os.ReadFile(path.Join(manifestsPath, filename))
	if err != nil {
		return filename, err
	}
	return string(bytesFile), err
}

func (r FileSystem) getFilenameFromDir(path string) (filename string, err error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return filename, err
	}
	for _, file := range files {
		if !file.IsDir() {
			filename = file.Name()
		}
	}
	if filename == "" {
		return "", errors.New("file not found")
	}

	return filename, nil
}

func (r FileSystem) Create(params *manifest.CreateParams) error {
	manifestPath := path.Join(r.basePath, params.Project, params.Tag)
	err := os.MkdirAll(manifestPath, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	perm := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	files, err := os.ReadDir(manifestPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		os.Remove(file.Name())
	}
	f, err := os.OpenFile(path.Join(manifestPath, params.Filename), perm, 0750)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewBuffer([]byte(params.Filename)))
	if err != nil {
		os.RemoveAll(path.Join(r.basePath, params.Project))
		return err
	}
	return nil
}

func (r FileSystem) Delete(params *manifest.BaseParams) error {
	var searchPath = path.Join(r.basePath, params.Project)
	files, err := os.ReadDir(searchPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			recSearchPath := path.Join(searchPath, file.Name())
			sFiles, err := os.ReadDir(recSearchPath)
			if err != nil {
				return err
			}
			for _, file := range sFiles {
				if file.Name() == params.Tag {
					os.Remove(path.Join(recSearchPath, params.Tag))
				}
			}
		}
	}
	if err != nil {
		return err
	}
	return nil
}
