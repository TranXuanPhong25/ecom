package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	jwtServiceClient pb.JWTServiceClient
	jwtServiceConn   *grpc.ClientConn
	jwtOnce          sync.Once
)

func InitJWTServiceClient(addr string) {
	jwtOnce.Do(func() {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		var err error
		jwtServiceConn, err = grpc.NewClient(addr, opts...)
		if err != nil {
			panic(fmt.Sprintf("failed to dial: %v", err))
		}

		jwtServiceClient = pb.NewJWTServiceClient(jwtServiceConn)
		log.Infof("Successfully connected to JWT service at %s", addr)
	})
}
func CloseJWTServiceConnection() {
	if jwtServiceConn != nil {
		if err := jwtServiceConn.Close(); err != nil {
			log.Errorf("Error closing jwt service connection: %v", err)
		} else {
			log.Info("Jwt service connection closed successfully")
		}
	}
}
func createToken(userId string) (string, error) {
	if jwtServiceClient == nil {
		return "", fmt.Errorf("jwt service client not initialized")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()
	r, err := jwtServiceClient.CreateToken(ctx, &pb.CreateTokenRequest{UserId: userId})
	if err != nil {
		log.Infof("could not create token: %v", err)
		return "", err
	}
	return r.GetToken(), nil
}
