package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/TranXuanPhong25/ecom/user/models"
	pb "github.com/TranXuanPhong25/ecom/user/proto"
	"github.com/TranXuanPhong25/ecom/user/repositories"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

//func LoadEnv() {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//	err = godotenv.Load(".env.example")
//	if err != nil {
//		log.Fatal("Error loading .env.example file")
//	}
//
//}

func (s *UserService) CreateUserWithEmailAndPassword(ctx context.Context, in *pb.Credentials) (*emptypb.Empty, error) {
	newUser := models.User{
		Email:    in.Email,
		Password: in.Password,
	}
	err := repositories.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) GetUserByEmailAndPassword(ctx context.Context, in *pb.Credentials) (*pb.User, error) {
	var targetUser models.User
	err := repositories.DB.
		Where("email = ? AND password = ?", in.Email, in.Password).
		First(&targetUser).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &pb.User{
		UserId: targetUser.ID.String(),
		Email:  in.Email,
	}, nil
}

func (s *UserService) DeleteUserById(ctx context.Context, in *pb.UserId) (*emptypb.Empty, error) {
	uid, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid uuid: %w", err)
	}

	err = repositories.DB.Delete(&models.User{}, uid).Error
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func RegisterService(server *grpc.Server) {
	pb.RegisterUserServiceServer(server, &UserService{})
	log.Println("UserService registered")
}
