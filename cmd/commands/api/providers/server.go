package providers

import (
	"github.com/MigAru/poseidon/internal/base"
	"github.com/MigAru/poseidon/internal/blob"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/gin"
	"github.com/MigAru/poseidon/internal/manifest"
	"github.com/MigAru/poseidon/internal/ping"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/sirupsen/logrus"
)

func ServerProvider(
	cfg *config.Config,
	log *logrus.Logger,
	pingController *ping.PingController,
	blobController *blob.Controller,
	baseController *base.Controller,
	manifestController *manifest.Controller,
) http.Server {
	return gin.NewServer(cfg, log, pingController, blobController, baseController, manifestController)
}
