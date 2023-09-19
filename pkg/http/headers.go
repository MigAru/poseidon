package http

import "fmt"

type DefaultParams struct {
	ContentLength int
	Range         bool
	StartRange    int
	EndRange      int
	Location      string
	ContentType   string
	UUID          string
	Digest        string
}

func NewRegisryHeadersParams() *DefaultParams {
	return &DefaultParams{}
}

func (p *DefaultParams) WithContentLength(length int) *DefaultParams {
	p.ContentLength = length
	return p
}

func (p *DefaultParams) WithRange(start, end int) *DefaultParams {
	p.Range = true
	p.StartRange = start
	p.EndRange = end
	return p
}

func (p *DefaultParams) WithLocation(location string) *DefaultParams {
	p.Location = location
	return p
}

func (p *DefaultParams) WithDigest(digest string) *DefaultParams {
	p.Digest = digest
	return p
}

func (p *DefaultParams) WithUUID(uuid string) *DefaultParams {
	p.UUID = uuid
	return p
}

func (p *DefaultParams) WithContentType(content string) *DefaultParams {
	p.ContentType = content
	return p
}

func CreateRegistryHeaders(params *DefaultParams) []Header {
	headers := []Header{
		{
			Key:   "Content-Length",
			Value: fmt.Sprintf("%d", params.ContentLength),
		},
	}

	if params.Range {
		headers = append(headers, Header{
			Key:   "Range",
			Value: fmt.Sprintf("%d-%d", params.StartRange, params.EndRange),
		})
	}

	if params.Location != "" {
		headers = append(headers, Header{
			Key:   "Location",
			Value: params.Location,
		})
	}

	if params.UUID != "" {
		headers = append(headers, Header{
			Key:   "Docker-Upload-UUID",
			Value: params.UUID,
		})
	}

	if params.Digest != "" {
		headers = append(headers, Header{
			Key:   "Docker-Content-Digest",
			Value: params.Digest,
		})
	}

	if params.ContentType != "" {
		headers = append(headers, Header{
			Key:   "Content-Type",
			Value: params.ContentType,
		})
	}

	return headers
}
