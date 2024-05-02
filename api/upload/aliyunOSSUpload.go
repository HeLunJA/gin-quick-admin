package upload

import (
	"github.com/gin-gonic/gin"
	"gvaTemplate/model/system/request"
	"gvaTemplate/model/system/response"
	"gvaTemplate/utils/upload"
	"strconv"
	"time"
)

type FileUploadAndDownloadApi struct{}

// UploadFile
// @Tags      uploads
// @Summary   上传文件示例
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=nil,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	claims, flag := c.Get("claims")
	if !flag {
		response.NoAuth("Object not found", c)
		return
	}
	exClaims, ok := claims.(request.BaseClaims)
	if !ok {
		response.Fail("Invalid object type", c)
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.Fail("No file uploaded", c)
		return
	}
	defer file.Close()
	oss := upload.NewOSS()
	fileName := strconv.FormatUint(uint64(exClaims.UserId), 10) + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + header.Filename
	filePath, key, uploadErr := oss.UploadFile(file, fileName)
	if uploadErr != nil {
		response.Fail(uploadErr.Error(), c)
	}
	response.Ok(gin.H{"filePath": filePath, "key": key}, "上传成功", c)
}
