package di

import (
	"context"
	"log"
	"tiny_files/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitializeMinioClient() *minio.Client {
	// Initialize MinIO client
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create bucket if it doesn't exist
	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, config.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, config.BucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", config.BucketName)
		} else {
			log.Fatal(err)
		}
	}

	return minioClient
}
