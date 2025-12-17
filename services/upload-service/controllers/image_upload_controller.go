package controllers

import (
	"net/http"
	"strconv"

	"github.com/TranXuanPhong25/ecom/services/upload-service/models"
	"github.com/TranXuanPhong25/ecom/services/upload-service/services"
)

func GeneratePresignedURLUploadImage(c echo.Context) error {
	req := new(models.GeneratePresignedURLUploadImagePayload)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	presignedURL, err := services.GeneratePresignedURLUploadImage(req.FileName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"title": "Failed to generate presigned URL",
			"error": err.Error(),
		})
	}
	resource := services.TruncateUrl(presignedURL)

	return c.JSON(http.StatusOK, map[string]string{
		"presignedUrl": presignedURL.String(),
		"filename":     req.FileName,
		"fileSize":     string(req.FileSize),
		"httpMethod":   req.HttpMethod,
		"mimeType":     req.MimeType,
		"resource":     resource,
		"expiresIn":    strconv.FormatInt(services.PresignedURLExpiration.Milliseconds(), 10),
	})
}
