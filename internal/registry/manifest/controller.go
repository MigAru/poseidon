package manifets

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"poseidon/internal/interfaces/digest/digest"
	"poseidon/internal/interfaces/manifest"
	"poseidon/pkg/http"
	v2_2 "poseidon/pkg/registry/manifest/schema/v2.2"
	"strconv"
	"strings"
)

type Controller struct {
	log        *logrus.Logger
	repository manifest.Repository
	digest     digest.Repository
}

func NewController(log *logrus.Logger, repository manifest.Repository, digest digest.Repository) *Controller {
	return &Controller{
		log:        log,
		repository: repository,
		digest:     digest,
	}
}

func (c Controller) Get(ctx http.Context) (err error) {
	name := ctx.Param("name")
	var (
		filename  = ctx.Param("tag")
		fileBytes []byte
	)
	params := manifest.NewGetParams(name, filename)
	if !isFilename(filename) {
		filename, err = c.repository.Get(params)
		if err != nil {
			return
		}
	}
	fileBytes, err = c.digest.Get(name, filename)
	if err != nil {
		return
	}
	var manifest v2_2.Manifest
	json.Unmarshal(fileBytes, &manifest)
	ctx.SetHeader("Docker-Content-Digest", filename)
	ctx.SetHeader("Content-Type", manifest.MediaType)
	ctx.SetHeader("Content-Length", strconv.Itoa(manifest.GetLength()))
	ctx.JSON(200, &manifest)
	return nil
}

func isFilename(name string) bool {
	s := strings.Split(name, ":")
	if len(s) > 1 {
		return true
	}

	return false
}

func (c Controller) Create(ctx http.Context) error {

	b, err := io.ReadAll(ctx.Body())
	if err != nil {
		return err
	}
	hasher := sha256.New()
	hasher.Write(b)
	hash := fmt.Sprintf("sha256:%x", hasher.Sum(nil))

	params := manifest.NewCreateParams(ctx.Param("name"), ctx.Param("tag"))
	if err := c.repository.Create(params.WithFilename(hash).WithData(b)); err != nil {
		ctx.NoContent(400)
		return err
	}

	if err := c.digest.Create(ctx.Param("name"), hash, b); err != nil {
		ctx.NoContent(400)
		return err
	}
	location := "/v2/" + ctx.Param("name") + "/manifest/" + hash
	ctx.SetHeader("Location", location)
	ctx.SetHeader("Docker-Content-Digest", hash)
	ctx.NoContent(201)
	return nil
}

func (c Controller) Delete(ctx http.Context) (err error) {
	reference := ctx.Param("tag")
	if !isFilename(reference) {
		params := manifest.NewGetParams(ctx.Param("name"), reference)
		reference, err = c.repository.Get(params)
	}

	err = c.repository.Delete(manifest.NewBaseParams(ctx.Param("name"), reference))
	if err != nil {
		return
	}

	err = c.digest.Delete(ctx.Param("name"), reference)
	if err != nil {
		return
	}
	return
}
