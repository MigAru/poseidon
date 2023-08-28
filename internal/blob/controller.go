package blob

import (
	"context"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/internal/upload"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	httpInterface "net/http"
	"strconv"
	"strings"
)

type Controller struct {
	log       *logrus.Logger
	fs        *file_system.FS
	chunkSize int
	manager   *upload.Manager
}

func NewController(log *logrus.Logger, cfg *config.Config, fs *file_system.FS, manager *upload.Manager) *Controller {
	return &Controller{log: log, chunkSize: cfg.Upload.ChunkSize, fs: fs, manager: manager}
}

//TODO: подумать над тем чтобы вынести в отдельный контроллер загрузок

func (c Controller) GetUpload(ctx http.Context) error {

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

func (c Controller) CreateUpload(ctx http.Context) error {
	// создание загрузки
	//TODO: реализовать менеджер загрузок и после реализовать двойное наименование name/subName:tag
	name := ctx.Param("project")
	subName := ctx.Param("sub-name-project")
	projectName := name
	if subName != "" {
		projectName += "." + subName
	}

	totalSize, err := strconv.Atoi(ctx.Header("Content-Length"))
	if err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}
	id, err := c.manager.Create(context.Background(), projectName, totalSize)
	if err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}

	uploadURL := "/v2/" + strings.ReplaceAll(projectName, ".", "/") + "/blobs/uploads/" + id

	//TODO: сделать header builder
	headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, 0).WithUUID(id)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))

	ctx.NoContent(httpInterface.StatusAccepted)
	return nil
}

func (c Controller) Upload(ctx http.Context) error {
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

func (c Controller) buildStatusUpload(uploadedSize, totalSize int) int {
	if uploadedSize == totalSize {
		return httpInterface.StatusCreated
	}
	return httpInterface.StatusAccepted
}

func (c Controller) DeleteUpload(ctx http.Context) error {
	if err := c.manager.Delete(ctx.Param("uuid")); err != nil {
		ctx.JSON(httpInterface.StatusNotFound, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return err
	}
	return nil
}

func (c Controller) Get(ctx http.Context) error {
	//TODO: сделать кэш отдачи слоев в памяти(middleware)
	digest := ctx.Param("digest")
	data, err := c.fs.GetDigest(ctx.Param("project"), digest)
	if err != nil {
		ctx.JSON(404, errors.NewErrorResponse(errors.BlobUnknown))
		return err
	}

	headers := http.NewRegisryHeadersParams().WithContentLength(len(data)).WithDigest(digest)

	if ctx.Request().Method == httpInterface.MethodHead {
		//TODO: сделать header builder

		ctx.SetHeaders(http.CreateRegistryHeaders(headers))
		ctx.NoContent(httpInterface.StatusOK)
		return nil
	}

	//TODO: сделать header builder
	ctx.SetHeaders(http.CreateRegistryHeaders(headers.WithContentType(http.ContentOctetStream)))
	ctx.OctetStream(httpInterface.StatusOK, data)
	return nil
}
