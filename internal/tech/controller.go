package tech

import (
	"github.com/MigAru/poseidon/internal/database"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/sirupsen/logrus"
	http2 "net/http"
)

type Controller struct {
	db  database.DB
	log *logrus.Logger
}

func NewController(log *logrus.Logger, db database.DB) *Controller {
	return &Controller{
		db:  db,
		log: log,
	}
}

type ListTagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tech"`
}

func (c *Controller) ListTags(ctx http.Context) {
	repositoryName := http.GetProjectName(ctx)

	tags, err := c.db.GetTags(repositoryName)
	if err != nil {
		ctx.JSON(http2.StatusNotFound, errors.NewErrorResponse(errors.NameUnknown))
		return
	}

	ctx.JSON(http2.StatusOK, ListTagsResponse{Name: repositoryName, Tags: tags})
}

type CatalogResponse struct {
	Repositories []string `json:"repositories"`
}

func (c *Controller) CatalogRepositories(ctx http.Context) {
	repositories, err := c.db.GetRepositories()
	if err != nil {
		ctx.JSON(http2.StatusNotFound, errors.NewErrorResponse(errors.NameUnknown))
		return
	}

	ctx.JSON(http2.StatusOK, CatalogResponse{Repositories: repositories})
}
