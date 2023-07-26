package file_system

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path"
)

//TODO: разбить по файлам

type BaseParamsManifest struct {
	Project string
	Tag     string
}

func NewBaseParamsManifest(project, tag string) *BaseParamsManifest {
	return &BaseParamsManifest{Project: project, Tag: tag}
}

type GetParamsManifest struct {
	*BaseParamsManifest
	Filename string
}

func NewGetParamsManifest(project, tag string) *GetParamsManifest {
	return &GetParamsManifest{BaseParamsManifest: NewBaseParamsManifest(project, tag)}
}

func (p *GetParamsManifest) WithFilename(name string) *GetParamsManifest {
	p.Filename = name
	return p
}

type CreateParamsManifest struct {
	*BaseParamsManifest
	Filename string
	Data     []byte
}

func NewCreateParamsManifest(project, tag string) *CreateParamsManifest {
	return &CreateParamsManifest{BaseParamsManifest: NewBaseParamsManifest(project, tag)}
}

func (p *CreateParamsManifest) WithFilename(name string) *CreateParamsManifest {
	p.Filename = name
	return p
}

func (p *CreateParamsManifest) WithData(bytes []byte) *CreateParamsManifest {
	p.Data = bytes
	return p
}

func (f *FS) GetManifest(params *GetParamsManifest) (filename string, err error) {
	manifestsPath := path.Join(f.basePath, params.Project, params.Tag)
	if params.Filename == "" {
		filename, err = f.getFilenameFromDir(manifestsPath)
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

func (f *FS) getFilenameFromDir(path string) (filename string, err error) {
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

func (f *FS) CreateManifest(params *CreateParamsManifest) error {
	manifestPath := path.Join(f.basePath, params.Project, params.Tag)
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
		if err := os.Remove(file.Name()); err != nil {
			return err
		}
	}
	file, err := os.OpenFile(path.Join(manifestPath, params.Filename), perm, 0750)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, bytes.NewBuffer([]byte(params.Filename)))
	if err != nil {
		if err := os.RemoveAll(path.Join(f.basePath, params.Project)); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (f *FS) DeleteManifest(params *BaseParamsManifest) error {
	var searchPath = path.Join(f.basePath, params.Project)
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
					if err := os.Remove(path.Join(recSearchPath, params.Tag)); err != nil {
						return err
					}
				}
			}
		}
	}
	if err != nil {
		return err
	}
	return nil
}
