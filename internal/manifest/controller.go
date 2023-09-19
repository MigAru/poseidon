package manifest

import (
	"fmt"
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/MigAru/poseidon/pkg/registry/hasher/methods"
	"github.com/sirupsen/logrus"
	"io"
	http2 "net/http"
	"os"
	"strings"
)

type Controller struct {
	log      *logrus.Logger
	manifest *Manager
	uploads  *upload.Manager
}

//TODO: сделать обработку ошибок

func NewController(log *logrus.Logger, uploads *upload.Manager, manifest *Manager) *Controller {
	return &Controller{
		log:      log,
		manifest: manifest,
		uploads:  uploads,
	}
}

func (c *Controller) Get(ctx http.Context) error {
	//TODO: сделать валидацию на project и reference
	var (
		project   = ctx.Param("project")
		reference = ctx.Param("reference")
	)

	manifest, digest, err := c.manifest.Get(project, reference)
	if os.IsNotExist(err) {
		ctx.JSON(http2.StatusNotFound, errors.NewErrorResponse(errors.NameUnknown))
		return err
	}
	if err != nil {
		ctx.JSON(http2.StatusBadRequest, errors.NewErrorResponse(errors.GetManifest))
		return err
	}

	headers := http.NewRegisryHeadersParams().
		WithDigest(digest).
		WithContentType(manifest.MediaType).
		WithContentLength(manifest.GetLength())

	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.JSON(200, manifest)
	return nil
}

func (c *Controller) isDigest(name string) bool {
	hashArray := strings.Split(name, ":")
	return len(hashArray) > 1
}

func (c *Controller) Create(ctx http.Context) error {
	var (
		project   = http.GetProjectName(ctx)
		reference = ctx.Param("reference")
	)

	data, err := io.ReadAll(ctx.Body())
	if err != nil {
		ctx.JSON(http2.StatusBadRequest, errors.NewErrorResponse(errors.GetManifest))
		return err
	}
	//TODO: сделать разбитие reference и парсинг метода
	hash, err := c.uploads.UploadManifest(project, reference, methods.SHA256, data)
	if err != nil {
		ctx.JSON(http2.StatusBadRequest, errors.NewErrorResponse(errors.CreateManifest))
		return err
	}

	location := fmt.Sprintf("/v2/%s/manifest/%s", strings.ReplaceAll(project, ".", "/"), hash)
	headers := http.NewRegisryHeadersParams().WithLocation(location).WithDigest(hash)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(201)
	return nil
}

func (c *Controller) Delete(ctx http.Context) error {
	project := ctx.Param("project")
	reference := ctx.Param("reference")

	if err := c.manifest.Delete(project, reference); err != nil {
		ctx.NoContent(http2.StatusInternalServerError)
		return err
	}
	return nil
}
