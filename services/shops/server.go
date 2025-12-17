package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/TranXuanPhong25/ecom/shops/configs"
	"github.com/TranXuanPhong25/ecom/shops/middlewares"
	"github.com/TranXuanPhong25/ecom/shops/models"
	"github.com/TranXuanPhong25/ecom/shops/repositories"
	"github.com/TranXuanPhong25/ecom/shops/routes"
	"github.com/TranXuanPhong25/ecom/shops/services"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func newEchoServer() *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.PrometheusMiddleware())
	routes.ShopsRoute(e)
	return e
}
func newRPCServer() (*grpc.Server, *net.Listener) {
	lis, err := net.Listen("tcp", ":"+configs.AppConfig.RPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	services.RegisterService(s)
	reflection.Register(s)
	return s, &lis
}

func main() {

	configs.LoadEnv()
	s, lis := newRPCServer()
	e := newEchoServer()
	repositories.ConnectDB()
	err := repositories.DB.AutoMigrate(&models.Shop{})
	if err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Printf("gRPC server listening on %v\n", configs.AppConfig.RPCPort)
		if err := s.Serve(*lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		log.Printf("HTTP server listening on %v\n", configs.AppConfig.ServerPort)
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutting down servers...\n")
	s.GracefulStop()
	e.Shutdown(context.Background())

	wg.Wait()
}
