package services

import (
	"context"
	"encoding/base64"
	"os"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"

	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
)

var (
	keyVault struct {
		privateKey []byte
	}
	SigningMethod = jwt.SigningMethodHS256
	ExpireTime    = time.Duration(36)
)

func LoadEnv() {
	secretKeyBase64 := os.Getenv("JWT_SECRET_KEY")
	if secretKeyBase64 == "" {
		log.Fatal("JWT_SECRET_KEY environment variable is not set")
	}
	decodedSecretKey, err := base64.StdEncoding.DecodeString(secretKeyBase64)
	if err != nil {
		log.Fatalf("Failed to decode JWT secret key: %v", err)
	}
	keyVault.privateKey = decodedSecretKey
}

type JWTService struct {
	pb.UnimplementedJWTServiceServer
}

func (s *JWTService) ValidateToken(ctx context.Context, in *pb.TokenRequest) (*pb.ValidationResponse, error) {
	response := &pb.ValidationResponse{
		UserId: "",
	}
	parsedToken, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		return keyVault.privateKey, nil
	}, jwt.WithLeeway(time.Minute))
	if err != nil {
		log.Printf("Failed to parse token: %v", err)
		return response, nil
	}

	if !parsedToken.Valid {
		log.Printf("Token is invalid")
		return response, nil
	}
	userId, err := parsedToken.Claims.GetSubject()
	if err != nil {
		log.Printf("Failed to get user id: %v", err)
	}
	return &pb.ValidationResponse{UserId: userId}, nil
}

func (s *JWTService) CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	unsignedToken := jwt.NewWithClaims(SigningMethod, jwt.MapClaims{
		"sub": in.UserId,
		"aud": in.Roles,
		"iss": "ecom-jwt-service",
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(ExpireTime * time.Minute)),
	})

	signedToken, err := unsignedToken.SignedString(keyVault.privateKey)

	return &pb.CreateTokenResponse{Token: signedToken}, err
}

func RegisterService(server *grpc.Server) {
	pb.RegisterJWTServiceServer(server, &JWTService{})
	log.Info("JWTService registered")
}
