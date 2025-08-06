/*
	A "hello world" plugin in Go,
	which reads a request header and sets a response header.
*/

package main

import (
	"sync"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	pb "github.com/TranXuanPhong25/ecom/jwt-service/proto"
	"google.golang.org/grpc"
)

var (
	Priority          = 1
	Version           = "0.2"
	JWTServiceAddr    = "jwt-service:50050"
	OPAServerRouteURL = "http://opa-server:8181/v1/data/route"
	client            pb.JWTServiceClient
	conn              *grpc.ClientConn
	once              sync.Once
)

type Config struct{}

func New() interface{} {
	return &Config{}
}

func main() {
	initClient(JWTServiceAddr)
	server.StartServer(New, Version, Priority)

}

func (conf Config) Access(kong *pdk.PDK) {
	token := getTokenFromHeaders(kong)
	validateToken(token, kong)
}

func (conf Config) Close() error {
	if conn != nil {
		return conn.Close()
	}
	return nil
}
