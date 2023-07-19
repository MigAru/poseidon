package manifest

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	digestInterface "poseidon/internal/interfaces/digest/digest"
	manifestInterface "poseidon/internal/interfaces/manifest"
	"poseidon/pkg/http"
	v2_2 "poseidon/pkg/registry/manifest/schema/v2.2"
	"strconv"
	"strings"
)

type Controller struct {
	log        *logrus.Logger
	repository manifestInterface.Repository
	digest     digestInterface.Repository
}

//TODO: сделать manifest manager
//TODO: сделать обработку ошибок

func NewController(log *logrus.Logger, repository manifestInterface.Repository, digest digestInterface.Repository) *Controller {
	return &Controller{
		log:        log,
		repository: repository,
		digest:     digest,
	}
}

func (c Controller) Get(ctx http.Context) (err error) {
	//TODO: сделать валидацию на project и reference
	var (
		project   = ctx.Param("project")
		filename  = ctx.Param("reference")
		fileBytes []byte
	)
	params := manifestInterface.NewGetParams(project, filename)
	if !c.isDigest(filename) {
		filename, err = c.repository.Get(params)
		if err != nil {
			return
		}
	}
	fileBytes, err = c.digest.Get(project, filename)
	if err != nil {
		return
	}
	//TODO: сделать универсальный unmarshaler для manifest v2 v1/oci/manifest list v2
	var manifest v2_2.Manifest
	json.Unmarshal(fileBytes, &manifest)
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

func (c Controller) Create(ctx http.Context) error {
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

	params := manifestInterface.NewCreateParams(project, reference)
	if err := c.repository.Create(params.WithFilename(hash).WithData(b)); err != nil {
		ctx.NoContent(400)
		return err
	}

	if err := c.digest.Create(project, hash, b); err != nil {
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

func (c Controller) Delete(ctx http.Context) (err error) {
	//TODO: сделать после менеджера загрузки
	reference := ctx.Param("tag")
	project := ctx.Param("project")
	if !c.isDigest(reference) {
		params := manifestInterface.NewGetParams(project, reference)
		reference, err = c.repository.Get(params)
		if err != nil {
			return err
		}
	}

	err = c.repository.Delete(manifestInterface.NewBaseParams(project, reference))
	if err != nil {
		return
	}

	err = c.digest.Delete(ctx.Param("name"), reference)
	if err != nil {
		return
	}
	return
}
