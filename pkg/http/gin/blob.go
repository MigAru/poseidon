package gin

import (
	"github.com/gin-gonic/gin"
	"poseidon/internal/registry/blob"
)

func (s *Server) registerBlobController(group *gin.RouterGroup, pattern string, controller *blob.Controller) {
	uploadPattern := pattern + ":uuid"
	// init upload
	group.POST(pattern, func(ctx *gin.Context) {
		controller.CreateUpload(WrapContext(ctx))
	})

	//getting info upload
	group.GET(uploadPattern, func(ctx *gin.Context) {
		controller.GetUpload(WrapContext(ctx))
	})
	group.HEAD(uploadPattern, func(ctx *gin.Context) {
		controller.GetUpload(WrapContext(ctx))
	})

	//uploading blob
	group.PATCH(uploadPattern, func(ctx *gin.Context) {
		controller.Upload(WrapContext(ctx))
	})
	group.PUT(uploadPattern, func(ctx *gin.Context) {
		controller.Upload(WrapContext(ctx))
	})

	// deleting blob
	group.DELETE(uploadPattern, func(ctx *gin.Context) {
		controller.Delete(WrapContext(ctx))
	})
}
