package http

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Context interface {
	JSON(code int, i interface{})
	Bind(i any) error
	FormFile(name string) (*multipart.FileHeader, error)
	Body() io.ReadCloser
	NoContent(code int)
	Param(string) string
	QueryParams() url.Values
	QueryParam(name string) string
	OriginalURL() string
	Request() *http.Request
	Redirect(location string, status int) error
	Header(key string) string
	SetHeader(key string, val string)
	SetHeaders([]Header)
	OctetStream(code int, data []byte)
}

type Handler interface {
	Handle(ctx Context) error
}

type HandlerFunc func(ctx Context) error

func (h HandlerFunc) Handle(ctx Context) error {
	return h(ctx)
}
