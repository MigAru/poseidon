package registry_api

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"poseidon/internal/config"
	"poseidon/internal/http/gin"
	manifest2 "poseidon/internal/interfaces/manifest"
	"poseidon/internal/ping"
	"poseidon/internal/registry/base"
	"poseidon/internal/registry/blob"
	manifest "poseidon/internal/registry/manifest"
	"poseidon/pkg/http"
)

var httpSet = wire.NewSet(
	ping.NewPingController,
	base.NewController,
	blob.NewController,
	manifest.NewController,
	wire.Bind(new(manifest2.Controller), new(*manifest.Controller)),
	ServerProvider,
)

func ServerProvider(
	cfg *config.Config,
	log *logrus.Logger,
	pingController *ping.PingController,
	blobController *blob.Controller,
	baseController *base.Controller,
	manifestController manifest2.Controller,
) http.HttpServer {
	return gin.NewServer(cfg, log, pingController, blobController, baseController, manifestController)
}
