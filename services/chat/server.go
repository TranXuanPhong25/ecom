package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TranXuanPhong25/ecom/services/chat/configs"
	"github.com/TranXuanPhong25/ecom/services/chat/controllers"
	"github.com/TranXuanPhong25/ecom/services/chat/repositories"
	"github.com/TranXuanPhong25/ecom/services/chat/routes"
	"github.com/TranXuanPhong25/ecom/services/chat/services"
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
	chatController := controllers.NewChatController(chatService)

	routes.RegisterChatRoutes(e, chatController)
	routes.RegisterHealthRoute(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		e.Logger.Fatal(e.Start(":" + configs.AppConfig.ServerPort))
	}()

	<-quit
}
