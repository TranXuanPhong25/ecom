package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/TranXuanPhong25/ecom/services/auth/models"
	pb "github.com/TranXuanPhong25/ecom/services/users/proto"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	usersServiceClient pb.UsersServiceClient
	usersServiceConn   *grpc.ClientConn
	usersOnce          sync.Once
)

func InitUsersServiceClient(addr string) {
	usersOnce.Do(func() {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		var err error
		usersServiceConn, err = grpc.NewClient(addr, opts...)
		if err != nil {
			log.Errorf(fmt.Sprintf("failed to dial: %v", err))
		}

		usersServiceClient = pb.NewUsersServiceClient(usersServiceConn)
		log.Infof("Successfully connected to users service at %s", addr)

	})
}
func CloseUsersServiceConnection() {
	if usersServiceConn != nil {
		if err := usersServiceConn.Close(); err != nil {
			log.Errorf("Error closing users service connection: %v", err)
		} else {
			log.Info("User service connection closed successfully")
		}
	}
}
func CreateUserWithEmailAndPassword(email, password string) error {
	if usersServiceClient == nil {
		return fmt.Errorf("users service client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()

	_, err := usersServiceClient.CreateUserWithEmailAndPassword(ctx, &pb.Credentials{Email: email, Password: password})
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmailAndPassword(email, password string) (*models.UserInfo, error) {
	if usersServiceClient == nil {
		return &models.UserInfo{}, fmt.Errorf("users service client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5000)*time.Millisecond)
	defer cancel()

	r, err := usersServiceClient.GetUserByEmailAndPassword(ctx, &pb.Credentials{
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

func GetUserById(id string) (*models.UserInfo, error) {
	if usersServiceClient == nil {
		return &models.UserInfo{}, fmt.Errorf("users service client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5000)*time.Millisecond)
	defer cancel()

	r, err := usersServiceClient.GetUserById(ctx, &pb.UserId{UserId: id})

	if err != nil {
		return &models.UserInfo{}, err
	}

	if r.GetUserId() == "" {
		return &models.UserInfo{}, fmt.Errorf("user not found")
	}

	return &models.UserInfo{UserId: r.GetUserId(), Email: r.GetEmail()}, nil
}
