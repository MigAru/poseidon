package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"io"
	httpInterface "net/http"
)

func (c *Controller) PutUpload(ctx http.Context) error {
	var (
		project = http.GetProjectName(ctx)
		uuid    = ctx.Param("uuid")
		digest  = ctx.QueryParam("digest")
	)

	if _, err := c.uploads.Get(uuid); err != nil {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return err
	}

	tx, err := c.db.NewTx(ctx.Request().Context())
	if err != nil {
		ctx.NoContent(httpInterface.StatusInternalServerError)
		return err
	}
	defer tx.Rollback()

	buffer, err := io.ReadAll(ctx.Body())
	if err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}

	written, err := c.uploads.Done(uuid, digest, buffer)
	if err != nil {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.DigestInvalid))
		return err
	}
	if err := c.db.IndexingDigest(tx, digest); err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}

	uploadURL := "/v2/" + project + "/blobs/uploads/" + uuid

	headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, written).WithUUID(uuid)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(201)
	return tx.Commit()
}
