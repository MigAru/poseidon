package file_system

type FS struct {
	basePath string
}

func New(template string) *FS {
	return &FS{basePath: template}
}
