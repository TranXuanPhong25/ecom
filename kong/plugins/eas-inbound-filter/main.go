/*
	A "hello world" plugin in Go,
	which reads a request header and sets a response header.
*/

package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	server.StartServer(New, Version, Priority)
	<-quit
	close()
}

var (
	Priority = 1
	Version  = "0.2"

	client pb.JWTServiceClient
	conn   *grpc.ClientConn
	once   sync.Once
)

type Config struct {
	Timeout        int
	JWTServiceAddr string
}

func New() interface{} {
	initClient("jwt-service:50051")
	return &Config{
		JWTServiceAddr: "jwt-service:50051", // Default address
		Timeout:        5000,                // Default timeout (5 seconds)
	}
}

func initClient(addr string) pb.JWTServiceClient {
	once.Do(func() {
		//follow https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.NewClient(addr, opts...)
		if err != nil {
			panic(fmt.Sprintf("failed to dial: %v", err))
		}

		client = pb.NewJWTServiceClient(conn)
	})
	return client
}

func (conf Config) Access(kong *pdk.PDK) {
	token := getTokenFromHeaders(kong)
	if token == "" {
		return
	}
	validateToken(token, kong)
}

func getTokenFromHeaders(kong *pdk.PDK) string {
	token := ""
	authorizationHeader, getAuthorizationError := kong.Request.GetHeader("Authorization")
	if getAuthorizationError == nil {
		token = getTokenFromAuthorizationHeader(authorizationHeader)
		if token != "" {
			return token
		}
	}

	cookieHeader, getCookieError := kong.Request.GetHeader("Cookie")
	if getCookieError == nil {
		token = getTokenFromCookieHeader(cookieHeader)
		if token != "" {
			return token
		}
	}
	return token
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

func validateToken(token string, kong *pdk.PDK) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000)*time.Millisecond)
	defer cancel()
	r, err := client.ValidateToken(ctx, &pb.TokenRequest{Token: token})
	if err != nil {
		kong.Log.Err("could not validate token: %v", err)
		return
	}

	if !r.GetValid() {
		kong.Response.Exit(401, []byte(`{"error":"Unauthorized"}`), map[string][]string{
			"Content-Type": {"application/json"},
		})
	}
}

// Thêm hàm này để clean up kết nối khi plugin bị hủy
func close() {
	if conn != nil {
		conn.Close()
	}
}
