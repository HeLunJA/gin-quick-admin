package upload

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gvaTemplate/global"
	"mime/multipart"
)

type AliyunOSS struct{}

type OSS interface {
	UploadFile(file multipart.File, fileName string) (string, string, error)
	DeleteFile(key string) error
}

func (*AliyunOSS) UploadFile(file multipart.File, fileName string) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}
	yunFileTmpPath := global.GT_CONFIG.AliyunOSS.BasePath + "/" + fileName

	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, file)
	if err != nil {
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return "https://" + global.GT_CONFIG.AliyunOSS.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket()
	if err != nil {
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(global.GT_CONFIG.AliyunOSS.Endpoint, global.GT_CONFIG.AliyunOSS.AccessKeyId, global.GT_CONFIG.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(global.GT_CONFIG.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func NewOSS() OSS {
	return &AliyunOSS{}
}
