package gin

import (
	"github.com/MigAru/poseidon/internal/manifest"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerManifestController(group *gin.RouterGroup, pattern string, controller *manifest.Controller) {
	group.PUT(pattern, func(ctx *gin.Context) {
		if err := controller.Upload(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	group.GET(pattern, func(ctx *gin.Context) {
		if err := controller.Get(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	group.HEAD(pattern, func(ctx *gin.Context) {
		if err := controller.Get(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	group.DELETE(pattern, func(ctx *gin.Context) {
		if err := controller.Delete(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
}
