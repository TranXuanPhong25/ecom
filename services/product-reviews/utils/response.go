package utils

import (
	"github.com/TranXuanPhong25/ecom/services/product-reviews/dtos"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SuccessResponse sends a success JSON response
func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, dtos.SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error JSON response
func ErrorResponse(c echo.Context, statusCode int, errorMsg string, message string) error {
	return c.JSON(statusCode, dtos.ErrorResponse{
		Error:   errorMsg,
		Message: message,
	})
}

// BadRequestError sends a 400 Bad Request error
func BadRequestError(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusBadRequest, "Bad Request", message)
}

// NotFoundError sends a 404 Not Found error
func NotFoundError(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusNotFound, "Not Found", message)
}

// ForbiddenError sends a 403 Forbidden error
func ForbiddenError(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusForbidden, "Forbidden", message)
}

// InternalServerError sends a 500 Internal Server Error
func InternalServerError(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", message)
}
