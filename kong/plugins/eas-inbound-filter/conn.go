package main

import (
	"fmt"

	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initClient(addr string) pb.JWTServiceClient {
	once.Do(func() {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.NewClient(addr, opts...)
		if err != nil {
			fmt.Printf("failed to dial: %v", err)
		}

		client = pb.NewJWTServiceClient(conn)
	})
	return client
}
