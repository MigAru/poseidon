package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	httpInterface "net/http"
)

func (c *Controller) GetUpload(ctx http.Context) error {

	var (
		uuid = ctx.Param("uuid")
	)

	blob, ok := c.manager.Get(uuid)
	if !ok {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return nil
	}

	uploaded := blob.UploadedBytes
	if uploaded > 0 {
		uploaded -= 1
	}

	headers := http.NewRegisryHeadersParams().WithRange(0, uploaded)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(httpInterface.StatusNoContent)
	return nil
}
