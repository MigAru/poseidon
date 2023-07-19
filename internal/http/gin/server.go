package gin

import (
	"context"
	"net/http"
	"poseidon/internal/config"
	"poseidon/internal/interfaces/blob"
	"poseidon/internal/interfaces/manifest"
	"poseidon/internal/ping"
	"poseidon/internal/registry/base"
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
	blobController blob.Controller,
	baseController *base.Controller,
	manifestController manifest.Controller,
) *Server {
	server := &Server{
		log:                   log,
		mainController:        gin.Default(),
		shutdownTimeoutSecond: time.Duration(cfg.Server.TimeoutGracefullShutdown),
		port:                  cfg.Server.Port,
	}

	server.registerControllers(pingController, blobController, baseController, manifestController)

	return server
}
func (s *Server) registerControllers(
	ping *ping.PingController,
	blobController blob.Controller,
	baseController *base.Controller,
	manifestController manifest.Controller,
) {
	s.registerPingController(ping)
	g := s.mainController.Group("/v2/")
	g.GET("", func(ctx *gin.Context) {
		if err := baseController.V2(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	g.HEAD(":name/blobs/:digest", func(ctx *gin.Context) {
		if err := blobController.Get(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	g.GET(":name/blobs/:digest", func(ctx *gin.Context) {
		if err := blobController.Get(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	s.registerManifestController(g, ":name/manifests/:tag", manifestController)
	s.registerUploadController(g, ":name/blobs/uploads/", blobController)
}

func (s *Server) registerPingController(controller *ping.PingController) {
	s.mainController.GET("/ping", func(ctx *gin.Context) {
		controller.Ping(WrapContext(ctx))
	})
}

func (s *Server) Run() {
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
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeoutSecond*time.Second)
	defer cancel()
	err := s.server.Shutdown(ctx)
	if err != nil {
		s.log.Error(err)
		return
	}
}