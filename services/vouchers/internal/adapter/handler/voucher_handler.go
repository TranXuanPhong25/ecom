package handler

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/port/in"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/validators"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// VoucherHandler handles HTTP requests for voucher operations
type VoucherHandler struct {
	service   in.VoucherServicePort
	validator *validator.Validate
}

// NewVoucherHandler creates a new instance of VoucherHandler
func NewVoucherHandler(service in.VoucherServicePort) *VoucherHandler {
	return &VoucherHandler{
		service:   service,
		validator: validators.NewValidator(),
	}
}

// CreateVoucher handles POST /vouchers
func (h *VoucherHandler) CreateVoucher(c echo.Context) error {
	var request dto.CreateVoucherRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	if err := h.validator.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	voucher, err := h.service.CreateVoucher(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    dto.ToVoucherResponse(voucher),
	})
}

// GetVoucherByCode handles GET /vouchers/code/:code
func (h *VoucherHandler) GetVoucherByCode(c echo.Context) error {
	code := c.Param("code")

	voucher, err := h.service.GetVoucherByCode(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    dto.ToVoucherResponse(voucher),
	})
}

// GetVoucherByID handles GET /vouchers/:id
func (h *VoucherHandler) GetVoucherByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid voucher ID",
		})
	}

	voucher, err := h.service.GetVoucherByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    dto.ToVoucherResponse(voucher),
	})
}

// ListVouchers handles GET /vouchers
func (h *VoucherHandler) ListVouchers(c echo.Context) error {
	var request dto.ListVouchersRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid query parameters",
		})
	}

	vouchers, total, err := h.service.ListVouchers(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Convert to response DTOs
	voucherResponses := make([]dto.VoucherResponse, len(vouchers))
	for i, v := range vouchers {
		voucherResponses[i] = dto.ToVoucherResponse(&v)
	}

	response := dto.ListVouchersResponse{
		Vouchers: voucherResponses,
		Total:    total,
		Page:     request.Page,
		Limit:    request.Limit,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    response,
	})
}

// UpdateVoucher handles PUT /vouchers/:id
func (h *VoucherHandler) UpdateVoucher(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid voucher ID",
		})
	}

	var request dto.UpdateVoucherRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	if err := h.validator.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	voucher, err := h.service.UpdateVoucher(c.Request().Context(), uint(id), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    dto.ToVoucherResponse(voucher),
	})
}

// DeleteVoucher handles DELETE /vouchers/:id
func (h *VoucherHandler) DeleteVoucher(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid voucher ID",
		})
	}

	if err := h.service.DeleteVoucher(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Voucher deleted successfully",
	})
}

// HealthCheck handles GET /health
func (h *VoucherHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "healthy",
	})
}
