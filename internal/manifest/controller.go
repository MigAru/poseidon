package manifest

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/pkg/http"
	v2_2 "github.com/MigAru/poseidon/pkg/registry/manifest/schema/v2.2"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"strings"
)

type Controller struct {
	log *logrus.Logger
	fs  *file_system.FS
}

//TODO: сделать manifest manager
//TODO: сделать обработку ошибок

func NewController(log *logrus.Logger, fs *file_system.FS) *Controller {
	return &Controller{
		log: log,
		fs:  fs,
	}
}

func (c *Controller) Get(ctx http.Context) (err error) {
	//TODO: сделать валидацию на project и reference
	var (
		project   = ctx.Param("project")
		filename  = ctx.Param("reference")
		fileBytes []byte
	)
	params := file_system.NewGetParamsManifest(project, filename)
	if !c.isDigest(filename) {
		filename, err = c.fs.GetManifest(params)
		if err != nil {
			return
		}
	}
	fileBytes, err = c.fs.GetDigest(project, filename)
	if err != nil {
		return
	}
	//TODO: сделать универсальный unmarshaler для manifest v2 v1/oci/manifest list v2
	var manifest v2_2.Manifest
	if err := json.Unmarshal(fileBytes, &manifest); err != nil {
		return err
	}
	//TODO: сделать header builder
	ctx.SetHeader("Docker-Content-Digest", filename)
	ctx.SetHeader("Content-Type", manifest.MediaType)
	ctx.SetHeader("Content-Length", strconv.Itoa(manifest.GetLength()))
	ctx.JSON(200, &manifest)
	return nil
}

func (c *Controller) isDigest(name string) bool {
	hashArray := strings.Split(name, ":")
	return len(hashArray) > 1
}

func (c *Controller) Create(ctx http.Context) error {
	//TODO: сделать валидацию на пустые проекты и референсы
	var (
		project   = ctx.Param("project")
		reference = ctx.Param("reference")
	)
	b, err := io.ReadAll(ctx.Body())
	if err != nil {
		return err
	}

	hasher := sha256.New()
	hasher.Write(b)
	hash := fmt.Sprintf("sha256:%x", hasher.Sum(nil))

	params := file_system.NewCreateParamsManifest(project, reference)
	if err := c.fs.CreateManifest(params.WithFilename(hash).WithData(b)); err != nil {
		ctx.NoContent(400)
		return err
	}

	if err := c.fs.CreateDigest(project, hash, b); err != nil {
		ctx.NoContent(400)
		return err
	}
	//TODO: сделать перенаправление на blob endpoints
	//TODO: сделать header builder
	location := "/v2/" + ctx.Param("name") + "/manifest/" + hash
	ctx.SetHeader("Location", location)
	ctx.SetHeader("Docker-Content-Digest", hash)
	ctx.NoContent(201)
	return nil
}

func (c *Controller) Delete(ctx http.Context) (err error) {
	//TODO: сделать после менеджера загрузки
	reference := ctx.Param("tag")
	project := ctx.Param("project")
	if !c.isDigest(reference) {
		params := file_system.NewGetParamsManifest(project, reference)
		reference, err = c.fs.GetManifest(params)
		if err != nil {
			return err
		}
	}

	err = c.fs.DeleteManifest(file_system.NewBaseParamsManifest(project, reference))
	if err != nil {
		return
	}

	err = c.fs.DeleteDigest(ctx.Param("name"), reference)
	if err != nil {
		return
	}
	return
}
