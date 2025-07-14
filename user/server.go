package main

import (
	"github.com/TranXuanPhong25/ecom/user/models"
	"github.com/TranXuanPhong25/ecom/user/repositories"
	"github.com/TranXuanPhong25/ecom/user/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	RpcPort = ":50052" // gRPC server port
)

func main() {
	//services.LoadEnv()
	runServer()
}
func runServer() {
	lis, err := net.Listen("tcp", RpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	repositories.ConnectDB()
	err = repositories.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	services.RegisterService(s)
	reflection.Register(s)
	log.Printf("gRPC server listening on %v\n", RpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
