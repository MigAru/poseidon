package manifest

import (
	"github.com/MigAru/poseidon/pkg/http"
	registryErrors "github.com/MigAru/poseidon/pkg/registry/errors"
	http2 "net/http"
)

func (c *Controller) Delete(ctx http.Context) error {
	ctx.JSON(http2.StatusNotFound, registryErrors.NewErrorResponse(registryErrors.Unsupported))
	return nil
}
