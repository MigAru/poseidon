package manifest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MigAru/poseidon/pkg/http"
	registryErrors "github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/MigAru/poseidon/pkg/registry/hasher/methods"
	v2_2 "github.com/MigAru/poseidon/pkg/registry/manifest/schema/v2.2"
	"io"
	http2 "net/http"
	"strings"
)

func (c *Controller) Upload(ctx http.Context) error {
	var (
		project   = http.GetProjectName(ctx)
		reference = ctx.Param("reference")
	)

	data, err := io.ReadAll(ctx.Body())
	if err != nil {
		ctx.JSON(http2.StatusBadRequest, registryErrors.NewErrorResponse(registryErrors.GetManifest))
		return err
	}

	var manifest v2_2.Manifest

	if err := json.Unmarshal(data, manifest); err != nil {
		ctx.JSON(http2.StatusInternalServerError, registryErrors.NewErrorResponse(registryErrors.ManifestInvalid))
		return err
	}

	tx, err := c.db.NewTx(ctx.Request().Context())
	if err != nil {
		ctx.NoContent(http2.StatusInternalServerError)
		return err
	}
	defer tx.Rollback()

	//TODO: сделать разбитие reference и парсинг метода
	//TODO: перевести на fs тк хранение текущего манифеста находится на бд

	hasher, err := c.hr.Build(methods.SHA256, data)
	if err != nil {
		return err
	}
	hash := fmt.Sprintf("%s:%x", methods.SHA256, hasher.Sum(nil))

	if err := c.fs.CreateDigest(hash, data); err != nil {
		ctx.JSON(http2.StatusBadRequest, registryErrors.NewErrorResponse(registryErrors.CreateManifest))
		return err
	}

	if err := c.createOrUpdateRepository(tx, project, reference, hash, manifest.GetLayersNames()); err != nil {
		ctx.NoContent(http2.StatusInternalServerError)
		return err
	}

	if err := tx.Commit(); err != nil {
		ctx.NoContent(http2.StatusInternalServerError)
		return err
	}

	location := fmt.Sprintf("/v2/%s/manifest/%s", strings.ReplaceAll(project, ".", "/"), hash)
	headers := http.NewRegisryHeadersParams().WithLocation(location).WithDigest(hash)
	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.NoContent(http2.StatusCreated)

	return nil
}

func (c *Controller) createOrUpdateRepository(tx *sql.Tx, project, tag, digest string, layers []string) error {
	_, err := c.db.GetRepository(tx, project, tag)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if errors.Is(err, sql.ErrNoRows) {
		if err := c.db.CreateRepository(tx, project, tag, digest); err != nil {
			return err
		}
	}

	if err := c.db.UpdateDigestRepository(tx, project, tag, digest); err != nil {
		return err
	}

	return nil
}
