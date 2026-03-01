package routes

import (
	"github.com/TranXuanPhong25/ecom/services/notifications/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterNotificationRoutes(e *echo.Echo, ctl controllers.INotificationController) {
	e.GET("/api/notifications/mine", ctl.GetMyNotifications)
	e.GET("/api/notifications/mine/unread-count", ctl.GetUnreadCount)
	e.POST("/api/notifications", ctl.CreateNotification)
	e.PUT("/api/notifications/mine/:id/read", ctl.MarkAsRead)
	e.PUT("/api/notifications/mine/read-all", ctl.MarkAllAsRead)
	e.DELETE("/api/notifications/mine/:id", ctl.DeleteNotification)
}
