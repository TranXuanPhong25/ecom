package services

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"time"
)

var (
	MinIOClient            *minio.Client
	ProductImageBucketName = "product-images"
	PresignedURLExpiration = 15 * time.Minute
)

func InitMinIOClient() {
	endpoint := "minio:9000"             // MinIO server endpoint
	accessKeyID := "myuseraccesskey"     // MinIO access key
	secretAccessKey := "myusersecretkey" // MinIO secret key
	useSSL := false                      // Set to true if using HTTPS

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	MinIOClient = client
	if err != nil {
		log.Errorf("%v\n", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()
	listBuckets, err := MinIOClient.ListBuckets(ctx)
	log.Printf("%v\n", listBuckets)
}

func GeneratePresignedURLUploadImage(objectName string) (string, error) {
	if MinIOClient == nil {
		log.Errorf("MinIO client is not initialized\n")
		return "", nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()
	presignedURL, err := MinIOClient.PresignedPutObject(ctx, ProductImageBucketName, objectName, PresignedURLExpiration)
	if err != nil {
		log.Errorf("%v\n", err)
		return "", err
	}
	// For local development, we need to set the host to localhost:9000 - this is because MinIO runs on localhost in a Docker container
	presignedURL.Host = "localhost:9000"
	return presignedURL.String(), nil
}
