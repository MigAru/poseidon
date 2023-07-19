package gin

import (
	"github.com/gin-gonic/gin"
	"poseidon/internal/interfaces/blob"
)

func (s *Server) registerBlobController(group *gin.RouterGroup, pattern string, controller *blob.Controller) {

}
