package gin

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	http2 "poseidon/pkg/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	gctx gin.Context
}

func WrapContext(ctx *gin.Context) http2.Context {
	return &Context{gctx: *ctx}
}

func (c *Context) JSON(code int, i any) {
	c.gctx.JSON(code, i)
}

func (c *Context) Body() io.ReadCloser {
	return c.gctx.Request.Body
}

func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	return c.gctx.FormFile(name)
}

func (c *Context) Bind(i any) error {
	return c.gctx.Bind(&i)
}

func (c *Context) NoContent(code int) error {
	c.gctx.AbortWithStatus(code)
	return nil
}

func (c *Context) Param(name string) string {
	return c.gctx.Param(name)
}

func (c *Context) QueryParams() url.Values {
	return c.gctx.Request.URL.Query()
}

func (c *Context) QueryParam(name string) string {
	return c.gctx.Query(name)
}

func (c *Context) OriginalURL() string {
	return c.gctx.Request.RequestURI
}

func (c *Context) Request() *http.Request {
	return c.gctx.Request
}

func (c *Context) Redirect(location string, status int) error {
	c.gctx.Redirect(status, location)
	return nil
}

func (c *Context) Header(key string, def string) string {
	return c.gctx.GetHeader(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.gctx.Header(key, value)
}

func (c *Context) OctetStream(code int, data []byte) {
	c.gctx.Data(code, "application/octet-stream", data)
}
