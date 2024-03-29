package base

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/sirupsen/logrus"
	http2 "net/http"
)

type Controller struct {
	log *logrus.Logger
}

func NewController(log *logrus.Logger) *Controller {
	return &Controller{log: log}
}

func (c Controller) V2(ctx http.Context) error {
	ctx.NoContent(http2.StatusOK)
	return nil
}

func (c Controller) Catalog(ctx http.Context) error {
	//TODO: после реализации всех манифестов и менеджеров - сделать листинг каталога
	ctx.NoContent(http2.StatusNotFound)
	return nil
}
