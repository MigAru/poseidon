package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
	httpInterface "net/http"
	"strings"
)

func (c *Controller) CreateUpload(ctx http.Context) error {
	// создание загрузки
	project := http.GetProjectName(ctx)

	id, err := c.uploads.Create()
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
