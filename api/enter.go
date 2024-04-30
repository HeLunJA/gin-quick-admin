package api

import (
	"gvaTemplate/api/system"
	"gvaTemplate/api/upload"
)

type ApiGroup struct {
	SystemApiGroup           system.ApiGroup
	FileUploadAndDownloadApi upload.FileUploadAndDownloadApi
}

var ApiGroupApp = new(ApiGroup)
