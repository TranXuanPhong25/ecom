package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TranXuanPhong25/ecom/services/chat/internal/adapter/in/http/handler"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/adapter/in/ws"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/adapter/out/repositories"
	"github.com/TranXuanPhong25/ecom/services/chat/internal/app/chat"
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
	chatService := chat.NewChatService(convRepo, msgRepo, lastReadRepo)
	chatHandler := handler.NewChatHandler(chatService)

	hub := chat.NewHub()
	go hub.Run()
	wsChatHandler := ws.NewWSChatHandler(hub)

	wsChatHandler.RegisterWSChatRoutes(e)
	chatHandler.RegisterChatRoutes(e)
	handler.RegisterHealthRoute(e)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	<-quit
}
