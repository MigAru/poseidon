package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	httpInterface "net/http"
)

func (c *Controller) Get(ctx http.Context) error {
	//TODO: сделать кэш отдачи слоев в памяти(middleware)
	digest := ctx.Param("digest")
	data, err := c.fs.GetDigest(ctx.Param("project"), digest)
	if err != nil {
		ctx.JSON(404, errors.NewErrorResponse(errors.BlobUnknown))
		return err
	}

	headers := http.NewRegisryHeadersParams().WithContentLength(len(data)).WithDigest(digest)

	if ctx.Request().Method == httpInterface.MethodHead {
		ctx.SetHeaders(http.CreateRegistryHeaders(headers))
		ctx.NoContent(httpInterface.StatusOK)
		return nil
	}

	ctx.SetHeaders(http.CreateRegistryHeaders(headers.WithContentType(http.ContentOctetStream)))
	ctx.OctetStream(httpInterface.StatusOK, data)
	return nil
}
