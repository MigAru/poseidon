package gin

import (
	"github.com/MigAru/poseidon/internal/blob"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerBlobController(group *gin.RouterGroup, pattern string, controller *blob.Controller) {
	group.HEAD(pattern, func(ctx *gin.Context) {
		if err := controller.Get(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	group.GET(pattern, func(ctx *gin.Context) {
		if err := controller.Get(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	group.DELETE(pattern, func(ctx *gin.Context) {
		controller.Delete(WrapContext(ctx))
	})
}
