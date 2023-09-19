package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	httpInterface "net/http"
)

func (c *Controller) DeleteUpload(ctx http.Context) error {
	if err := c.manager.DeleteUpload(ctx.Param("uuid")); err != nil {
		ctx.JSON(httpInterface.StatusNotFound, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return err
	}

	ctx.NoContent(httpInterface.StatusNoContent)
	return nil
}
