package manifest

import (
	"encoding/json"
	"github.com/MigAru/poseidon/pkg/http"
	registryErrors "github.com/MigAru/poseidon/pkg/registry/errors"
	v2_2 "github.com/MigAru/poseidon/pkg/registry/manifest/schema/v2.2"
	http2 "net/http"
	"os"
)

func (c *Controller) Get(ctx http.Context) error {
	var (
		project   = http.GetProjectName(ctx)
		reference = ctx.Param("reference")
		manifest  v2_2.Manifest
	)
	repository, err := c.db.GetRepository(nil, project, reference)
	if err != nil {
		ctx.JSON(http2.StatusNotFound, registryErrors.NewErrorResponse(registryErrors.NameUnknown))
		return err
	}

	fileBytes, err := c.fs.GetDigest(project, repository.Digest)
	if os.IsNotExist(err) {
		ctx.JSON(http2.StatusNotFound, registryErrors.NewErrorResponse(registryErrors.NameUnknown))
		return err
	}
	if err != nil {
		ctx.JSON(http2.StatusBadRequest, registryErrors.NewErrorResponse(registryErrors.GetManifest))
		return err
	}

	//TODO: сделать универсальный unmarshaler для manifest v2 v1/oci/manifest list v2
	if err := json.Unmarshal(fileBytes, &manifest); err != nil {
		ctx.JSON(http2.StatusBadRequest, registryErrors.NewErrorResponse(registryErrors.ManifestInvalid))
		return err
	}

	headers := http.NewRegisryHeadersParams().
		WithDigest(repository.Digest).
		WithContentType(manifest.MediaType).
		WithContentLength(manifest.GetLength())

	ctx.SetHeaders(http.CreateRegistryHeaders(headers))
	ctx.JSON(200, manifest)
	return nil
}
