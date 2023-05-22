package registry_api

import (
	"github.com/sirupsen/logrus"
	"poseidon/internal/config"
	"poseidon/internal/ping"
	"poseidon/internal/registry/base"
	"poseidon/internal/registry/blob"
	"poseidon/pkg/http"
	"poseidon/pkg/http/gin"

	"github.com/google/wire"
)

var httpSet = wire.NewSet(
	ping.NewPingController,
	base.NewController,
	blob.NewController,
	ServerProvider,
)

func ServerProvider(cfg *config.Config, log *logrus.Logger, pingController *ping.PingController, blobController *blob.Controller, baseController *base.Controller) http.HttpServer {
	return gin.NewServer(cfg, log, pingController, blobController, baseController)
}
