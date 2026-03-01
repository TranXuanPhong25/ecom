package controllers

import (
	"net/http"

	"github.com/TranXuanPhong25/ecom/services/notifications/dtos"
	"github.com/TranXuanPhong25/ecom/services/notifications/services"
	"github.com/TranXuanPhong25/ecom/services/notifications/utils"
	"github.com/labstack/echo/v4"
)

type INotificationController interface {
	GetMyNotifications(c echo.Context) error
	GetUnreadCount(c echo.Context) error
	CreateNotification(c echo.Context) error
	MarkAsRead(c echo.Context) error
	MarkAllAsRead(c echo.Context) error
	DeleteNotification(c echo.Context) error
}

type NotificationController struct {
	service services.INotificationService
}

func NewNotificationController(service services.INotificationService) INotificationController {
	return &NotificationController{service: service}
}

func getUserID(c echo.Context) (string, bool) {
	userID := c.Request().Header.Get("X-User-Id")
	return userID, userID != ""
}

func (ctl *NotificationController) GetMyNotifications(c echo.Context) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	notifications, err := ctl.service.GetMyNotifications(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, notifications)
}

func (ctl *NotificationController) GetUnreadCount(c echo.Context) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	count, err := ctl.service.GetUnreadCount(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int{"unreadCount": count})
}

func (ctl *NotificationController) CreateNotification(c echo.Context) error {
	req := new(dtos.CreateNotificationPayload)
	if err := utils.ValidateRequestStructure(c, req); err != nil {
		return err
	}
	created, err := ctl.service.CreateNotification(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (ctl *NotificationController) MarkAsRead(c echo.Context) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	id := c.Param("id")
	if err := ctl.service.MarkAsRead(id, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Notification marked as read"})
}

func (ctl *NotificationController) MarkAllAsRead(c echo.Context) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	if err := ctl.service.MarkAllAsRead(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "All notifications marked as read"})
}

func (ctl *NotificationController) DeleteNotification(c echo.Context) error {
	userID, ok := getUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}
	id := c.Param("id")
	if err := ctl.service.DeleteNotification(id, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
