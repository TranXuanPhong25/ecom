package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/TranXuanPhong25/ecom/auth/models"
	pb "github.com/TranXuanPhong25/ecom/users/proto"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userServiceClient pb.UserServiceClient
	userServiceConn   *grpc.ClientConn
	userOnce          sync.Once
)

func InitUserServiceClient(addr string) {
	userOnce.Do(func() {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		var err error
		userServiceConn, err = grpc.NewClient(addr, opts...)
		if err != nil {
			log.Errorf(fmt.Sprintf("failed to dial: %v", err))
		}

		userServiceClient = pb.NewUserServiceClient(userServiceConn)
		log.Infof("Successfully connected to user service at %s", addr)

	})
}
func CloseUserServiceConnection() {
	if userServiceConn != nil {
		if err := userServiceConn.Close(); err != nil {
			log.Errorf("Error closing user service connection: %v", err)
		} else {
			log.Info("User service connection closed successfully")
		}
	}
}
func CreateUserWithEmailAndPassword(email, password string) error {
	if userServiceClient == nil {
		return fmt.Errorf("user service client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()

	_, err := userServiceClient.CreateUserWithEmailAndPassword(ctx, &pb.Credentials{Email: email, Password: password})
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmailAndPassword(email, password string) (*models.UserInfo, error) {
	if userServiceClient == nil {
		return &models.UserInfo{}, fmt.Errorf("user service client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5000)*time.Millisecond)
	defer cancel()

	r, err := userServiceClient.GetUserByEmailAndPassword(ctx, &pb.Credentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return &models.UserInfo{}, err
	}

	if r.GetEmail() == "" {
		return &models.UserInfo{}, fmt.Errorf("user not found")
	}
	if r.GetUserId() == "" {
		return &models.UserInfo{}, fmt.Errorf("wrong password")
	}

	return &models.UserInfo{UserId: r.GetUserId(), Email: r.GetEmail()}, nil
}
