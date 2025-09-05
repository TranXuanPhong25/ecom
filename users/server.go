package main

import (
	"log"
	"net"

	"github.com/TranXuanPhong25/ecom/users/configs"
	"github.com/TranXuanPhong25/ecom/users/models"
	"github.com/TranXuanPhong25/ecom/users/repositories"
	"github.com/TranXuanPhong25/ecom/users/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configs.LoadEnv()
	runServer()
}
func runServer() {

	lis, err := net.Listen("tcp", configs.AppConfig.RpcPort)
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
	log.Printf("gRPC server listening on %v\n", configs.AppConfig.RpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
