package gin

import (
	"github.com/MigAru/poseidon/internal/blob"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerUploadController(group *gin.RouterGroup, pattern string, controller *blob.Controller) {
	uploadPattern := pattern + ":uuid"

	// init uploads
	group.POST(pattern, func(ctx *gin.Context) {
		if err := controller.CreateUpload(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})

	//getting info uploads
	group.GET(uploadPattern, func(ctx *gin.Context) {
		if err := controller.GetUpload(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	group.HEAD(uploadPattern, func(ctx *gin.Context) {
		if err := controller.GetUpload(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})

	//uploading blob
	group.PATCH(uploadPattern, func(ctx *gin.Context) {
		if err := controller.Upload(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
	group.PUT(uploadPattern, func(ctx *gin.Context) {
		if err := controller.Upload(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})

	// deleting blob
	group.DELETE(uploadPattern, func(ctx *gin.Context) {
		if err := controller.DeleteUpload(WrapContext(ctx)); err != nil {
			s.log.Error(err.Error())
		}
	})
}
