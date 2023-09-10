package blob

import (
	"context"
	"github.com/MigAru/poseidon/pkg/http"
	httpInterface "net/http"
	"strconv"
	"strings"
)

func (c *Controller) CreateUpload(ctx http.Context) error {
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

	headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, 0).WithUUID(id)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(httpInterface.StatusAccepted)
	return nil
}
