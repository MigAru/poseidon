package ping

import (
	"github.com/MigAru/poseidon/pkg/http"
)

type PingController struct {
}

func NewController() *PingController {
	return &PingController{}
}

func (c *PingController) Ping(ctx http.Context) {
	ctx.JSON(200, struct{ Message string }{Message: "pong"})
}
