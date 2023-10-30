package manifest

import (
	"context"
	"github.com/MigAru/poseidon/pkg/http"
	registryErrors "github.com/MigAru/poseidon/pkg/registry/errors"
	http2 "net/http"
)

func (c *Controller) Delete(ctx http.Context) error {
	project, reference := ctx.Param("project"), ctx.Param("reference")

	tx, err := c.db.NewTx(context.Background())
	if err != nil {
		ctx.NoContent(http2.StatusInternalServerError)
		return err
	}
	defer tx.Rollback()

	repository, err := c.db.GetRepository(tx, project, reference)
	if err != nil {
		ctx.JSON(http2.StatusNotFound, registryErrors.NewErrorResponse(registryErrors.NameUnknown))
		return err
	}

	if err := c.db.MarkDeleteRepository(tx, repository.ID); err != nil {
		ctx.NoContent(http2.StatusInternalServerError)
		return err
	}

	return tx.Commit()
}
