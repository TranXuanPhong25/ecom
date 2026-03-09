package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/adapter/in/http/handler"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/adapter/out/repositories"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/app/services"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/infras/configs"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	repositories.InitDBConnection()

	e := echo.New()

	convRepo := repositories.NewConversationRepository()
	msgRepo := repositories.NewMessageRepository()
	lastReadRepo := repositories.NewLastReadRepository()
	chatService := services.NewChatService(convRepo, msgRepo, lastReadRepo)
	chatHandler := handler.NewChatHandler(chatService)

	chatHandler.RegisterChatRoutes(e)
	handler.RegisterHealthRoute(e)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	<-quit
}
