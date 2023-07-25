package gin

import (
	"context"
	"github.com/MigAru/poseidon/internal/blob"
	"github.com/MigAru/poseidon/internal/config"
	"github.com/MigAru/poseidon/internal/interfaces/manifest"
	"github.com/MigAru/poseidon/internal/ping"
	"github.com/MigAru/poseidon/internal/registry/base"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log                   *logrus.Logger
	mainController        *gin.Engine
	shutdownTimeoutSecond time.Duration
	port                  string

	server *http.Server
}

func NewServer(
	cfg *config.Config,
	log *logrus.Logger,
	pingController *ping.PingController,
	blobController *blob.Controller,
	baseController *base.Controller,
	manifestController manifest.Controller,
) *Server {
	server := &Server{
		log:                   log,
		mainController:        gin.Default(),
		shutdownTimeoutSecond: time.Duration(cfg.Server.TimeoutGracefullShutdown) * time.Second,
		port:                  cfg.Server.Port,
	}

	server.registerControllers(pingController, blobController, baseController, manifestController)

	return server
}
func (s *Server) registerControllers(
	ping *ping.PingController,
	blobController *blob.Controller,
	baseController *base.Controller,
	manifestController manifest.Controller,
) {
	s.registerPingController(ping)

	APIv2Group := s.mainController.Group("/v2/")
	APIv2Group.GET("", func(ctx *gin.Context) {
		if err := baseController.V2(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	s.registerManifestController(APIv2Group, ":project/manifests/:reference", manifestController)
	s.registerBlobController(APIv2Group, ":project/blobs/:digest", blobController)
	s.registerUploadController(APIv2Group, ":project/blobs/uploads/", blobController)
}

func (s *Server) registerPingController(controller *ping.PingController) {
	s.mainController.GET("/ping", func(ctx *gin.Context) {
		controller.Ping(WrapContext(ctx))
	})
}

func (s *Server) Run(ctx context.Context) {
	s.server = &http.Server{
		Addr:    s.port,
		Handler: s.mainController,
	}

	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			s.log.Error(err)
			return
		}
		s.shutdown(ctx)
	}()
}

func (s *Server) shutdown(ctx context.Context) {
	<-ctx.Done()
	shutdown, cancel := context.WithTimeout(ctx, s.shutdownTimeoutSecond*time.Second)
	defer cancel()
	err := s.server.Shutdown(shutdown)
	if err != nil {
		s.log.Error(err)
		return
	}
}
