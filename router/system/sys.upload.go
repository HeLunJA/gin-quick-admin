package system

import (
	"github.com/gin-gonic/gin"
	v "gvaTemplate/api"
)

type uploadRouter struct{}

func (s *uploadRouter) InitUploadRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	uploadRouter := Router.Group("file")
	UploadApi := v.ApiGroupApp.FileUploadAndDownloadApi
	{
		uploadRouter.POST("upload", UploadApi.UploadFile)
	}
	return uploadRouter
}
