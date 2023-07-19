package blob

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	http2 "net/http"
	blob2 "poseidon/internal/interfaces/blob"
	"poseidon/internal/interfaces/digest/digest"
	"poseidon/pkg/http"
	"poseidon/pkg/registry/errors"
	"strconv"
)

type Controller struct {
	log    *logrus.Logger
	blob   blob2.Repository
	digest digest.Repository
}

//TODO: сделать uploads manager
//TODO: сделать обработку ошибок

func NewController(log *logrus.Logger, blob blob2.Repository, digest digest.Repository) *Controller {
	return &Controller{log: log, blob: blob, digest: digest}
}

//TODO: подумать над тем чтобы вынести в отдельный контроллер загрузок

func (c Controller) GetUpload(ctx http.Context) error {
	//нужно при разбитии запроса
	return nil
}

func (c Controller) CreateUpload(ctx http.Context) error {
	// создание временных слоев
	//TODO: реализовать менеджер загрузок и после реализовать двойное наименование name/subName:tag
	name := ctx.Param("project")
	subName := ctx.Param("sub-name-project")
	locationName := "/" + name
	if subName != "" {
		locationName += "/" + subName
	}

	id := uuid.NewString()
	//TODO: сделать header builder
	ctx.SetHeader("Location", "/v2"+locationName+"/blobs/uploads/"+id)
	//сколько взял байтов 0-0 - взял все
	ctx.SetHeader("Range", "bytes=0-0")
	ctx.SetHeader("Content-Length", "0")
	//уникальный id загрузки(чтобы докер отслеживал загрузку)
	ctx.SetHeader("Docker-Upload-UUID", id)
	ctx.NoContent(http2.StatusAccepted)
	return nil
}

func (c Controller) Upload(ctx http.Context) error {
	//TODO: разнести upload на patch и put
	var (
		buffer  bytes.Buffer
		project = ctx.Param("project")
		uuid    = ctx.Param("uuid")
	)
	if ctx.QueryParam("digest") != "" {
		//для того чтобы создать постоянный слой в памяти
		if ctx.Header("Content-Length", "") == "0" {
			data, _ := c.blob.Get(uuid)
			c.digest.Create(project, ctx.QueryParam("digest"), data)

			ctx.SetHeader("Location", "/v2/"+project+"/blobs/uploads/"+uuid)
			ctx.SetHeader("Content-Length", "0")
			ctx.SetHeader("Range", "0-"+strconv.Itoa(len(data)-1))
			ctx.SetHeader("Docker-Upload-UUID", uuid)

			ctx.NoContent(201)
			return nil
		}
	}
	//загрузка временного слоя
	_, err := io.Copy(&buffer, ctx.Body())
	if err != nil {
		ctx.JSON(http2.StatusBadRequest, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return err
	}
	if err := c.blob.Create(uuid, buffer.Bytes()); err != nil {
		return err
	}

	//TODO: сделать header builder и location builder
	ctx.SetHeader("Location", "/v2/"+project+"/blobs/uploads/"+uuid)
	ctx.SetHeader("Content-Length", "0")
	ctx.SetHeader("Range", "0-"+strconv.Itoa(buffer.Len()-1))
	ctx.SetHeader("Docker-Upload-UUID", uuid)
	ctx.NoContent(201)
	return nil
}

func (c Controller) DeleteUpload(ctx http.Context) error {
	//TODO: реализовать когда будет реализован менеджер загрузок
	return nil
}

func (c Controller) Get(ctx http.Context) error {
	//TODO: сделать кэш отдачи слоев в памяти(middleware)
	digest := ctx.Param("digest")
	data, err := c.digest.Get(ctx.Param("project"), digest)
	if err != nil {
		//TODO: сделать response builder
		ctx.JSON(404, errors.NewErrorResponse(errors.BlobUnknown))
		return err
	}
	if ctx.Request().Method == http2.MethodHead {
		//TODO: сделать header builder
		ctx.SetHeader("Content-Length", strconv.Itoa(len(data)))
		ctx.SetHeader("Docker-Content-Digest", digest)
		ctx.NoContent(http2.StatusOK)
		return nil
	}

	//TODO: сделать header builder
	ctx.SetHeader("Content-Length", strconv.Itoa(len(data)))
	ctx.SetHeader("Docker-Content-Digest", digest)
	ctx.SetHeader("Content-Type", "application/octet-stream")
	ctx.OctetStream(http2.StatusOK, data)
	return nil
}
