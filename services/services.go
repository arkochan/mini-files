// services.go
package services

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
)

func GenerateUploadURL(minioClient *minio.Client, bucketName, filename string) (string, error) {
	// Generate presigned URL for upload (PUT)
	url, err := minioClient.PresignedPutObject(context.Background(), bucketName, filename, time.Hour)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func GenerateDownloadURL(minioClient *minio.Client, bucketName, filename string) (string, error) {
	// Generate presigned URL for download (GET)
	url, err := minioClient.PresignedGetObject(context.Background(), bucketName, filename, time.Hour, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
