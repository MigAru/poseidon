package manifest

type Repository interface {
	Get(params *GetParams) (string, error)
	Create(params *CreateParams) error
	Delete(params *BaseParams) error
}

type BaseParams struct {
	Project string
	Tag     string
}

func NewBaseParams(project, tag string) *BaseParams {
	return &BaseParams{Project: project, Tag: tag}
}

type GetParams struct {
	*BaseParams
	Filename string
}

func NewGetParams(project, tag string) *GetParams {
	return &GetParams{BaseParams: NewBaseParams(project, tag)}
}

func (p *GetParams) WithFilename(name string) *GetParams {
	p.Filename = name
	return p
}

type CreateParams struct {
	*BaseParams
	Filename string
	Data     []byte
}

func NewCreateParams(project, tag string) *CreateParams {
	return &CreateParams{BaseParams: NewBaseParams(project, tag)}
}

func (p *CreateParams) WithFilename(name string) *CreateParams {
	p.Filename = name
	return p
}

func (p *CreateParams) WithData(bytes []byte) *CreateParams {
	p.Data = bytes
	return p
}
