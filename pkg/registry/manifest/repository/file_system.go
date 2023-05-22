package repository

import "path"

type FileSystem struct {
	basePath string
}

func NewFileSystem(basePath string) *FileSystem {
	basePath = path.Join(basePath, "manifests")
	return &FileSystem{basePath: basePath}
}
