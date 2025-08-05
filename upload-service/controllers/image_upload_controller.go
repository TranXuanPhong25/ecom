package controllers

import (
	"github.com/TranXuanPhong25/ecom/upload-service/models"
	"github.com/TranXuanPhong25/ecom/upload-service/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GeneratePresignedURLUploadImage(c echo.Context) error {
	req := new(models.GeneratePresignedURLUploadImagePayload)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	presignedURL, err := services.GeneratePresignedURLUploadImage(req.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"title": "Failed to generate presigned URL",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"presigned_url": presignedURL,
		"filename":      req.Filename,
		"file_size":     req.FileSize,
		"http_method":   req.HttpMethod,
		"mime_type":     req.MimeType,
		"resource":      req.Resource,
		"expires_in":    services.PresignedURLExpiration.String(),
	})
}
