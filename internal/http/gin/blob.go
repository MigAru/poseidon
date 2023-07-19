package gin

import (
	"github.com/gin-gonic/gin"
	"poseidon/internal/interfaces/blob"
)

func (s *Server) registerBlobController(group *gin.RouterGroup, pattern string, controller blob.Controller) {
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
}
