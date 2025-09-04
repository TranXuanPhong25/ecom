package main

import (
	"fmt"

	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getJWTSvcClient() pb.JWTServiceClient {
	once.Do(func() {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		newConn, err := grpc.NewClient(JWTServiceAddr, opts...)
		if err != nil {
			fmt.Printf("failed to dial: %v", err)
		}
		conn = newConn
		JWTSvcClient = pb.NewJWTServiceClient(conn)
	})
	return JWTSvcClient
}
