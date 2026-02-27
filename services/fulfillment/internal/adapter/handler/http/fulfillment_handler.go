package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rengumin/fulfillment/internal/core/dto"
	"github.com/rengumin/fulfillment/internal/core/port"
)

type FulfillmentHandler struct {
	service port.FulfillmentService
}

func NewFulfillmentHandler(service port.FulfillmentService) *FulfillmentHandler {
	return &FulfillmentHandler{service: service}
}

func (h *FulfillmentHandler) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/fulfillment")
	api.POST("/pickup/schedule", h.SchedulePickup)
	api.POST("/pickup/confirm", h.MarkPickedUp)
	api.POST("/location/update", h.UpdateLocation)
	api.POST("/delivery/status", h.UpdateDeliveryStatus)
	api.GET("/tracking/:packageNumber", h.GetTracking)
	api.GET("/packages", h.ListPackages)
	api.GET("/packages/order/:orderId", h.GetPackageByOrderID)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
}

// SchedulePickup schedules a pickup from seller
// @Summary Schedule pickup
// @Description Schedule a pickup from seller warehouse
// @Tags fulfillment
// @Accept json
// @Produce json
// @Param request body dto.SchedulePickupRequest true "Pickup request"
// @Success 200 {object} dto.SchedulePickupResponse
// @Router /api/v1/fulfillment/pickup/schedule [post]
func (h *FulfillmentHandler) SchedulePickup(c echo.Context) error {
	var req dto.SchedulePickupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	resp, err := h.service.SchedulePickup(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FulfillmentHandler) MarkPickedUp(c echo.Context) error {
	var req dto.MarkPickedUpRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.MarkPickedUp(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Package marked as picked up"})
}

func (h *FulfillmentHandler) UpdateLocation(c echo.Context) error {
	var req dto.UpdateLocationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.UpdateLocation(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Location updated successfully"})
}

func (h *FulfillmentHandler) UpdateDeliveryStatus(c echo.Context) error {
	var req dto.UpdateDeliveryStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.UpdateDeliveryStatus(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Delivery status updated"})
}

func (h *FulfillmentHandler) GetTracking(c echo.Context) error {
	packageNumber := c.Param("packageNumber")

	tracking, err := h.service.GetPackageTracking(c.Request().Context(), packageNumber)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tracking)
}

func (h *FulfillmentHandler) ListPackages(c echo.Context) error {
	var query dto.ListPackagesQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 20
	}

	resp, err := h.service.ListPackages(c.Request().Context(), query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FulfillmentHandler) GetPackageByOrderID(c echo.Context) error {
	orderID, err := strconv.ParseInt(c.Param("orderId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order ID"})
	}

	pkg, err := h.service.GetPackageByOrderID(c.Request().Context(), orderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, pkg)
}
