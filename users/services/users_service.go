package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/TranXuanPhong25/ecom/users/models"
	pb "github.com/TranXuanPhong25/ecom/users/proto"
	"github.com/TranXuanPhong25/ecom/users/repositories"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UsersService struct {
	pb.UnimplementedUsersServiceServer
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

func (s *UsersService) CreateUserWithEmailAndPassword(_ context.Context, in *pb.Credentials) (*emptypb.Empty, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := models.User{
		Email:    in.Email,
		Password: string(hashedPassword),
	}
	err = repositories.DB.Create(&newUser).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") ||
			strings.Contains(err.Error(), "unique constraint") ||
			strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, status.Error(codes.AlreadyExists, "Email already exists")
		}

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UsersService) GetUserByEmailAndPassword(_ context.Context, in *pb.Credentials) (*pb.User, error) {

	var user models.User
	err := repositories.DB.
		Where("email = ?", in.Email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return &pb.User{Email: user.Email}, nil
	}

	return &pb.User{
		UserId: user.ID.String(),
		Email:  in.Email,
	}, nil
}

func (s *UsersService) DeleteUserById(_ context.Context, in *pb.UserId) (*emptypb.Empty, error) {
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
	pb.RegisterUsersServiceServer(server, &UsersService{})
	log.Info("UsersService registered")
}
