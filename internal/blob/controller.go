package blob

import (
	"bytes"
	"github.com/MigAru/poseidon/internal/file_system"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	httpInterface "net/http"
	"strconv"
)

type Controller struct {
	log *logrus.Logger
	fs  *file_system.FS
}

//TODO: сделать uploads manager
//TODO: сделать обработку ошибок

func NewController(log *logrus.Logger, fs *file_system.FS) *Controller {
	return &Controller{log: log, fs: fs}
}

//TODO: подумать над тем чтобы вынести в отдельный контроллер загрузок

func (c Controller) GetUpload(_ http.Context) error {
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
	ctx.NoContent(httpInterface.StatusAccepted)
	return nil
}

func (c Controller) Upload(ctx http.Context) error {
	//TODO: разнести upload на patch и put
	var (
		buffer  bytes.Buffer
		project = ctx.Param("project")
		UUID    = ctx.Param("uuid")
	)
	if ctx.QueryParam("digest") != "" {
		//для того чтобы создать постоянный слой в памяти
		if ctx.Header("Content-Length") == "0" {
			data, _ := c.fs.GetBlob(UUID)
			if err := c.fs.CreateDigest(project, ctx.QueryParam("digest"), data); err != nil {
				return err
			}

			ctx.SetHeader("Location", "/v2/"+project+"/blobs/uploads/"+UUID)
			ctx.SetHeader("Content-Length", "0")
			ctx.SetHeader("Range", "0-"+strconv.Itoa(len(data)-1))
			ctx.SetHeader("Docker-Upload-UUID", UUID)

			ctx.NoContent(201)
			return nil
		}
	}
	//загрузка временного слоя
	_, err := io.Copy(&buffer, ctx.Body())
	if err != nil {
		ctx.JSON(httpInterface.StatusBadRequest, errors.NewErrorResponse(errors.BlobUploadUnknown))
		return err
	}
	if err := c.fs.CreateBlob(UUID, buffer.Bytes()); err != nil {
		return err
	}

	//TODO: сделать header builder и location builder
	ctx.SetHeader("Location", "/v2/"+project+"/blobs/uploads/"+UUID)
	ctx.SetHeader("Content-Length", "0")
	ctx.SetHeader("Range", "0-"+strconv.Itoa(buffer.Len()-1))
	ctx.SetHeader("Docker-Upload-UUID", UUID)
	ctx.NoContent(201)
	return nil
}

func (c Controller) DeleteUpload(_ http.Context) error {
	//TODO: реализовать когда будет реализован менеджер загрузок
	return nil
}

func (c Controller) Get(ctx http.Context) error {
	//TODO: сделать кэш отдачи слоев в памяти(middleware)
	digest := ctx.Param("digest")
	data, err := c.fs.GetDigest(ctx.Param("project"), digest)
	if err != nil {
		//TODO: сделать response builder
		ctx.JSON(404, errors.NewErrorResponse(errors.BlobUnknown))
		return err
	}
	if ctx.Request().Method == httpInterface.MethodHead {
		//TODO: сделать header builder
		ctx.SetHeader("Content-Length", strconv.Itoa(len(data)))
		ctx.SetHeader("Docker-Content-Digest", digest)
		ctx.NoContent(httpInterface.StatusOK)
		return nil
	}

	//TODO: сделать header builder
	ctx.SetHeader("Content-Length", strconv.Itoa(len(data)))
	ctx.SetHeader("Docker-Content-Digest", digest)
	ctx.SetHeader("Content-Type", "application/octet-stream")
	ctx.OctetStream(httpInterface.StatusOK, data)
	return nil
}
