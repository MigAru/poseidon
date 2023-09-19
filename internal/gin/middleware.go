package gin

import (
	httpInterface "github.com/MigAru/poseidon/pkg/http"
	"github.com/MigAru/poseidon/pkg/registry/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) validateProjectNameMiddleware(ctx *gin.Context) {
	if err := httpInterface.ValidateProjectNameMiddleware(WrapContext(ctx)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrorResponse(errors.NameUnknown))
		return
	}
	ctx.Next()
}
