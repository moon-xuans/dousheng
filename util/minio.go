package util

import (
	"context"
	"dousheng/config"
	"errors"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	videoType = "video/mp4"
	imageType = "image/jpeg"
)

func UploadFile(objectName, fileType string, reader io.Reader) (string, error) {
	context := context.Background()
	// 创建一个MinIO客户端
	endpoint := fmt.Sprintf("%s:%d", config.Info.Server.IP, config.Info.MPort)
	accessKeyID := config.Info.AccessKeyId
	secretAccessKey := config.Info.SecretAccessKey
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV2(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return "", err
	}

	var contentType string
	bucketName := config.Info.Bucket
	if fileType == "video" {
		objectName = fmt.Sprintf("%s/%s", config.Info.VideoDir, objectName)
		contentType = videoType
	} else if fileType == "image" {
		objectName = fmt.Sprintf("%s/%s", config.Info.ImageDir, objectName)
		contentType = imageType
	} else {
		return "", errors.New("上传至minio的类型不支持")
	}

	// 上传文件到MinIO服务器
	uploadInfo, err := minioClient.PutObject(context, bucketName, objectName, reader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	return uploadInfo.Location, nil
}
