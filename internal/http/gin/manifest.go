package gin

import (
	"github.com/gin-gonic/gin"
	"poseidon/internal/interfaces/manifest"
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
