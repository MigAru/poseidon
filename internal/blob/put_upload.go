package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"io/ioutil"
	httpInterface "net/http"
)

func (c *Controller) PutUpload(ctx http.Context) error {
	var (
		project = http.GetProjectName(ctx)
		uuid    = ctx.Param("uuid")
		digest  = ctx.QueryParam("digest")
	)

	blob, ok := c.manager.Get(uuid)
	if !ok {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return nil
	}

	buffer, err := ioutil.ReadAll(ctx.Body())
	if err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}
	written, err := blob.Done(digest, buffer)
	if err != nil {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.DigestInvalid))
		return err
	}

	uploadURL := "/v2/" + project + "/blobs/upload/" + uuid

	headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, written).WithUUID(uuid)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(201)
	return nil
}
