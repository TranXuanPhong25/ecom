/*
	A "hello world" plugin in Go,
	which reads a request header and sets a response header.
*/

package main

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/TranXuanPhong25/jwt-service/proto"
)

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "0.1"
var Priority = 1

type Config struct {
	Timeout         int    `json:"timeout"`
	Message         string `json:"message"`
	AuthServiceAddr string `json:"auth_service_addr"` // Địa chỉ gRPC service
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	var token string
	authorizationHeader, getAuthorizationError := kong.Request.GetHeader("Authorization")
	if getAuthorizationError == nil {
		token = getTokenFromAuthorizationHeader(authorizationHeader)
	}
	cookieHeader, getCookieError := kong.Request.GetHeader("Cookie")
	if getCookieError == nil {
		token = getTokenFromCookieHeader(cookieHeader)
	}
	if token == "" {
		kong.Response.Exit(401, []byte("Unauthorized access"), map[string][]string{
			"Content-Type": {"application/json"},
		})
		return
	}

	//follow https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(conf.AuthServiceAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("fail to close connection: %v", err)
		}
	}(conn)

	client := pb.NewJWTServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.ValidateToken(ctx, &pb.TokenRequest{Token: token})
	if err != nil {
		log.Fatalf("could not validate token: %v", err)
	}
	log.Printf("the token is: %v", r.GetValid())
}

func getTokenFromAuthorizationHeader(authorizationHeader string) string {
	if len(authorizationHeader) > 7 && authorizationHeader[:7] == "Bearer " {
		return authorizationHeader[7:]
	}
	return ""
}
func getTokenFromCookieHeader(cookieHeader string) string {
	tokenField := []string{"token", "Token", "TOKEN"}
	// Parse the cookie header to extract the token
	cookieParts := bytes.Split([]byte(cookieHeader), []byte("; "))
	for _, part := range cookieParts {
		for _, field := range tokenField {
			if bytes.HasPrefix(part, []byte(field+"=")) {
				return string(part[len(field)+1:])
			}
		}
	}
	return ""
}
