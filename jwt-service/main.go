package main

import (
	"log"
	"net"

	services "github.com/TranXuanPhong25/ecom/jwt-service/services"
	"google.golang.org/grpc"
)

var (
	RpcPort = ":50051" // gRPC server port
)

func main() {
	services.LoadEnv()
	runServer()
}
func runServer() {
	lis, err := net.Listen("tcp", RpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	services.RegisterService(s)
	log.Printf("gRPC server listening on %v\n", RpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
