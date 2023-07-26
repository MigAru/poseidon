package providers

import (
	"github.com/MigAru/poseidon/internal/blob"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/http/gin"
	manifestInterface "github.com/MigAru/poseidon/internal/interfaces/manifest"
	"github.com/MigAru/poseidon/internal/ping"
	"github.com/MigAru/poseidon/internal/registry/base"
	"github.com/MigAru/poseidon/internal/registry/manifest"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var httpSet = wire.NewSet(
	ping.NewPingController,
	base.NewController,
	blob.NewController,
	manifest.NewController,
	wire.Bind(new(manifestInterface.Controller), new(*manifest.Controller)),
	ServerProvider,
)

func ServerProvider(
	cfg *config.Config,
	log *logrus.Logger,
	pingController *ping.PingController,
	blobController *blob.Controller,
	baseController *base.Controller,
	manifestController manifestInterface.Controller,
) http.Server {
	return gin.NewServer(cfg, log, pingController, blobController, baseController, manifestController)
}