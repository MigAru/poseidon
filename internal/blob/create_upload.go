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
	project := http.GetProjectName(ctx)

	totalSize, err := strconv.Atoi(ctx.Header("Content-Length"))
	if err == nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}
	id, err := c.manager.Create(context.Background(), project, totalSize)
	if err != nil {
		ctx.NoContent(httpInterface.StatusBadRequest)
		return err
	}

	uploadURL := "/v2/" + strings.ReplaceAll(project, ".", "/") + "/blobs/uploads/" + id

	headers := http.NewRegisryHeadersParams().WithLocation(uploadURL).WithRange(0, 0).WithUUID(id)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(httpInterface.StatusAccepted)
	return nil
}
