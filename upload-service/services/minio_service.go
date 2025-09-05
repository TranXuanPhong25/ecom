package services

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/TranXuanPhong25/ecom/upload-service/configs"
	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	MinIOClient            *minio.Client
	PresignedURLExpiration = 15 * time.Minute
)

func InitMinIOClient() {
	endpoint := configs.AppConfig.MinIOEndpoint                   // MinIO server endpoint
	accessKeyID := configs.AppConfig.MinIOAccessKey               // MinIO access key
	secretAccessKey := configs.AppConfig.MinIOSecretKey           // MinIO secret key
	useSSL, _ := strconv.ParseBool(configs.AppConfig.MinIOUseSSL) // Set to true if using HTTPS

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
	presignedURL, err := MinIOClient.PresignedPutObject(ctx, configs.AppConfig.MinIOBucketName, objectName, PresignedURLExpiration)
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
