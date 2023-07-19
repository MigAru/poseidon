package api

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"poseidon/internal/config"
	"poseidon/internal/http/gin"
	blobInterface "poseidon/internal/interfaces/blob"
	manifestInterface "poseidon/internal/interfaces/manifest"
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
	wire.Bind(new(blobInterface.Controller), new(*blob.Controller)),
	manifest.NewController,
	wire.Bind(new(manifestInterface.Controller), new(*manifest.Controller)),
	ServerProvider,
)

func ServerProvider(
	cfg *config.Config,
	log *logrus.Logger,
	pingController *ping.PingController,
	blobController blobInterface.Controller,
	baseController *base.Controller,
	manifestController manifestInterface.Controller,
) http.HttpServer {
	return gin.NewServer(cfg, log, pingController, blobController, baseController, manifestController)
}
