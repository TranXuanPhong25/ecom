package services

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"time"
)

var (
	MinIOClient            *minio.Client
	ProductImageBucketName = "product-images"
	PresignedURLExpiration = 15 * time.Minute
)

func InitMinIOClient() {
	endpoint := "minio:9000"       // MinIO server endpoint
	accessKeyID := "admin"         // MinIO access key
	secretAccessKey := "admin1234" // MinIO secret key
	useSSL := false                // Set to true if using HTTPS

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Errorf("%v\n", err)
		return
	}
	if client == nil {
		log.Error("MinIO client is nil")
		return
	}
	MinIOClient = client
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()
	listBuckets, err := MinIOClient.ListBuckets(ctx)
	log.Printf("%v", listBuckets)
}

func GeneratePresignedURLUploadImage(objectName string) (*url.URL, error) {
	if MinIOClient == nil {
		log.Errorf("MinIO client is not initialized\n")
		return &url.URL{}, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()
	presignedURL, err := MinIOClient.PresignedPutObject(ctx, ProductImageBucketName, objectName, PresignedURLExpiration)
	if err != nil {
		log.Errorf("%v\n", err)
		return &url.URL{}, err
	}
	// Modify host to point to minio proxy server
	presignedURL.Host = "localhost:9000"
	return presignedURL, nil
}

func TruncateUrl(url *url.URL) string {
	return url.Scheme + "://" + url.Host + url.Path
}
