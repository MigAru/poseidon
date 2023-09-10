package manifest

import (
	"fmt"
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/hasher/methods"
	"github.com/sirupsen/logrus"
	"io"
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
	if err != nil {
		return err
	}
	fmt.Println(manifest.GetLength())
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
		project   = ctx.Param("project")
		reference = ctx.Param("reference")
	)

	data, err := io.ReadAll(ctx.Body())
	if err != nil {
		return err
	}

	hash, err := c.uploads.UploadManifest(project, reference, methods.SHA256, data)
	if err != nil {
		return err
	}

	location := "/v2/" + ctx.Param("name") + "/manifest/" + hash
	headers := http.NewRegisryHeadersParams().WithLocation(location).WithDigest(hash)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(201)
	return nil
}

func (c *Controller) Delete(ctx http.Context) error {
	project := ctx.Param("project")
	reference := ctx.Param("reference")

	if err := c.manifest.Delete(project, reference); err != nil {
		return err
	}
	return nil
}
