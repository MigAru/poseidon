package blob

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	http2 "net/http"
	"poseidon/internal/registry/digest"
	"poseidon/pkg/http"
	"poseidon/pkg/registry"
	"strconv"
)

type Controller struct {
	log    *logrus.Logger
	blob   Repository
	digest digest.Repository
}

func NewController(log *logrus.Logger, blob Repository, digest digest.Repository) *Controller {
	return &Controller{log: log, blob: blob, digest: digest}
}

func (c *Controller) GetUpload(ctx http.Context) error {

	//if ctx.Request().Method == http2.MethodHead {
	//	ctx.JSON(http2.StatusNotFound, registry.ErrorResponse{Errors: []registry.Error{
	//		{
	//			Code:    "",
	//			Message: "Not Found layer",
	//			Detail:  "it's a new image because not exist",
	//		},
	//	}})
	//	return nil
	//}

	ctx.NoContent(http2.StatusOK)
	return nil
}

func (c *Controller) CreateUpload(ctx http.Context) error {

	name := ctx.Param("name")
	subName := ctx.Param("sub-name")
	locationName := "/" + name
	if subName != "" {
		locationName += "/" + subName
	}

	uuid := uuid.NewString()

	ctx.SetHeader("Location", "/v2"+locationName+"/blobs/uploads/"+uuid)
	ctx.SetHeader("Range", "bytes=0-0")
	ctx.SetHeader("Content-Length", "0")

	ctx.SetHeader("Docker-Upload-UUID", uuid)
	ctx.NoContent(http2.StatusAccepted)

	return nil
}

func (c *Controller) Upload(ctx http.Context) error {
	var buffer bytes.Buffer
	if ctx.QueryParam("digest") != "" {
		if ctx.Header("Content-Length", "") == "0" {
			fmt.Println(ctx.QueryParam("digest"))
			data, _ := c.blob.Get(ctx.Param("uuid"))
			c.digest.Create(ctx.Param("name"), ctx.QueryParam("digest"), data)

		}
	}

	_, err := io.Copy(&buffer, ctx.Body())
	if err != nil {
		ctx.JSON(http2.StatusBadRequest, registry.ErrorResponse{Errors: []registry.Error{
			{
				Code:    "",
				Message: "octet-stream error",
				Detail:  "an error occurred while copying | " + err.Error(),
			},
		}})

		return err
	}
	if err := c.blob.Create(ctx.Param("uuid"), buffer.Bytes()); err != nil {
		return err
	}
	ctx.SetHeader("Location", "/v2/"+ctx.Param("name")+"/blobs/uploads/"+ctx.Param("uuid"))
	ctx.SetHeader("Content-Length", "0")
	ctx.SetHeader("Range", "0-"+strconv.Itoa(buffer.Len()-1))
	ctx.SetHeader("Docker-Upload-UUID", ctx.Param("uiid"))
	ctx.NoContent(201)
	return nil
}

func (c *Controller) DeleteUpload(ctx http.Context) error {
	return nil
}

func (c *Controller) Get