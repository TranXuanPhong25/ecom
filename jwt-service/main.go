package main

import (
	"context"
	"log"
	"net"

	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
	"google.golang.org/grpc"
)

type JWTService struct {
	pb.UnimplementedJWTServiceServer
}

func (s *JWTService) ValidateToken(ctx context.Context, in *pb.TokenRequest) (*pb.ValidationResponse, error) {
	log.Printf("Received token: %s", in.Token)
	// Thêm logic xác thực token thực tế ở đây
	return &pb.ValidationResponse{Valid: true}, nil
}

var (
	RpcPort = ":8202" // gRPC server port
)

func main() {
	lis, err := net.Listen("tcp", RpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterJWTServiceServer(s, &JWTService{})

	log.Printf("gRPC server listening on %v\n", RpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	lis.Close()
	s.Stop()
}
