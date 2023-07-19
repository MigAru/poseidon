package gin

import (
	"github.com/MigAru/poseidon/internal/interfaces/manifest"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerManifestController(group *gin.RouterGroup, pattern string, controller manifest.Controller) {
	group.PUT(pattern, func(ctx *gin.Context) {
		if err := controller.Create(WrapContext(ctx)); err != nil {
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

}
