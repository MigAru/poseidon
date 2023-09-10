package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"io/ioutil"
	httpInterface "net/http"
)

func (c *Controller) Upload(ctx http.Context) error {
	//TODO: разнести upload на patch и put
	var (
		project = ctx.Param("project")
		uuid    = ctx.Param("uuid")
		digest  = ctx.QueryParam("digest")
	)

	blob, ok := c.manager.Get(uuid)
	if !ok {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return nil
	}

	defer ctx.Body().Close()
	if ctx.QueryParam("digest") != "" && ctx.Header("Content-Length") == "0" {
		//для того чтобы создать постоянный слой в памяти

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

	buffer, err := ioutil.ReadAll(ctx.Body())
	totalBytes := len(buffer)
	if err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}

	if blob.UploadedBytes <= 0 {
		blob.TotalSize = totalBytes
	}

	if err := c.manager.Update(uuid, buffer); err != nil {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.RangeInvalid))
		return err
	}

	uploadURL := "/v2/" + project + "/blobs/uploads/" + uuid
	//docker client is not support Range
	headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, len(buffer)-1).WithUUID(uuid)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))

	ctx.NoContent(c.buildStatusUpload(blob.UploadedBytes, totalBytes))
	return nil
}
