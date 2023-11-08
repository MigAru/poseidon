package blob

import (
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"io"
	httpInterface "net/http"
	"strings"
)

func (c *Controller) Upload(ctx http.Context) error {
	//TODO: разнести upload на patch и put
	var (
		project = http.GetProjectName(ctx)
		uuid    = ctx.Param("uuid")
		digest  = ctx.QueryParam("digest")
	)

	blobRaw, err := c.uploads.Get(uuid)
	if err != nil {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return err
	}

	defer ctx.Body().Close()
	if ctx.QueryParam("digest") != "" && ctx.Header("Content-Length") == "0" {
		//для того чтобы создать постоянный слой в памяти

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

		uploadURL := "/v2/" + strings.ReplaceAll(project, ".", "/") + "/blobs/upload/" + uuid

		headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, written).WithUUID(uuid)
		ctx.SetHeaders(http.CreateRegistryHeaders(headers))
		ctx.NoContent(201)
		return nil
	}

	buffer, err := io.ReadAll(ctx.Body())
	totalBytes := len(buffer)
	if err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}

	params := upload.NewUpdateParams(uuid).WithChunk(buffer)
	if err := c.uploads.Update(params); err != nil {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.RangeInvalid))
		return err
	}

	uploadURL := "/v2/" + strings.ReplaceAll(project, ".", "/") + "/blobs/uploads/" + uuid
	//docker client is not support Range
	headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, len(buffer)-1).WithUUID(uuid)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))

	ctx.NoContent(c.buildStatusUpload(len(blobRaw), totalBytes))
	return nil
}
