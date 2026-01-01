package services

import (
	"net/http"
	"time"

	"github.com/TranXuanPhong25/ecom/services/promotions/dtos"
	"github.com/TranXuanPhong25/ecom/services/promotions/models"
	"github.com/TranXuanPhong25/ecom/services/promotions/repositories"
	"github.com/labstack/echo/v4"
)

// ============ EVENT BANNER SERVICES ============

func CreateEventBanner(request *models.CreateEventBannerRequest) (*dtos.EventBannerDTO, *echo.HTTPError) {
	banner := &models.EventBanner{
		Title:       request.Title,
		Description: request.Description,
		ImageURL:    request.ImageURL,
		LinkURL:     request.LinkURL,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		Priority:    request.Priority,
		IsActive:    request.IsActive,
		EventType:   request.EventType,
		Position:    request.Position,
	}

	if banner.Position == "" {
		banner.Position = "main"
	}

	tx := repositories.DB.Create(banner)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while creating event banner")
	}

	return toEventBannerDTO(banner), nil
}

func GetEventBannerByID(id string) (*dtos.EventBannerDTO, *echo.HTTPError) {
	banner := &models.EventBanner{}
	tx := repositories.DB.First(banner, "id = ?", id)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Event banner not found")
	}
	return toEventBannerDTO(banner), nil
}

func GetAllEventBanners(activeOnly bool) ([]dtos.EventBannerDTO, *echo.HTTPError) {
	var banners []models.EventBanner
	query := repositories.DB.Order("priority DESC, created_at DESC")

	if activeOnly {
		now := time.Now()
		query = query.Where("is_active = ? AND start_time <= ? AND end_time >= ?", true, now, now)
	}

	tx := query.Find(&banners)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while fetching event banners")
	}

	result := make([]dtos.EventBannerDTO, 0)
	for _, banner := range banners {
		result = append(result, *toEventBannerDTO(&banner))
	}

	return result, nil
}

func UpdateEventBanner(request *models.UpdateEventBannerRequest) (*dtos.EventBannerDTO, *echo.HTTPError) {
	banner := &models.EventBanner{}
	tx := repositories.DB.First(banner, "id = ?", request.ID)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Event banner not found")
	}

	// Update only provided fields
	updates := make(map[string]interface{})
	if request.Title != "" {
		updates["title"] = request.Title
	}
	if request.Description != "" {
		updates["description"] = request.Description
	}
	if request.ImageURL != "" {
		updates["image_url"] = request.ImageURL
	}
	if request.LinkURL != "" {
		updates["link_url"] = request.LinkURL
	}
	if request.StartTime != nil {
		updates["start_time"] = request.StartTime
	}
	if request.EndTime != nil {
		updates["end_time"] = request.EndTime
	}
	if request.Priority != nil {
		updates["priority"] = *request.Priority
	}
	if request.IsActive != nil {
		updates["is_active"] = *request.IsActive
	}
	if request.EventType != "" {
		updates["event_type"] = request.EventType
	}
	if request.Position != "" {
		updates["position"] = request.Position
	}

	tx = repositories.DB.Model(banner).Updates(updates)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while updating event banner")
	}

	return toEventBannerDTO(banner), nil
}

func DeleteEventBanner(id string) *echo.HTTPError {
	tx := repositories.DB.Delete(&models.EventBanner{}, "id = ?", id)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while deleting event banner")
	}
	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Event banner not found")
	}
	return nil
}

// GetEventBannersByType - Lấy banners theo loại sự kiện
func GetEventBannersByType(eventType string, activeOnly bool) ([]dtos.EventBannerDTO, *echo.HTTPError) {
	var banners []models.EventBanner
	query := repositories.DB.Where("event_type = ?", eventType).Order("priority DESC, created_at DESC")

	if activeOnly {
		now := time.Now()
		query = query.Where("is_active = ? AND start_time <= ? AND end_time >= ?", true, now, now)
	}

	tx := query.Find(&banners)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while fetching event banners")
	}

	result := make([]dtos.EventBannerDTO, 0)
	for _, banner := range banners {
		result = append(result, *toEventBannerDTO(&banner))
	}
	return result, nil
}

// GetEventBannersByPosition - Lấy banners theo vị trí
func GetEventBannersByPosition(position string, activeOnly bool) ([]dtos.EventBannerDTO, *echo.HTTPError) {
	var banners []models.EventBanner
	query := repositories.DB.Where("position = ?", position).Order("priority DESC, created_at DESC")

	if activeOnly {
		now := time.Now()
		query = query.Where("is_active = ? AND start_time <= ? AND end_time >= ?", true, now, now)
	}

	tx := query.Find(&banners)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while fetching event banners")
	}

	result := make([]dtos.EventBannerDTO, 0)
	for _, banner := range banners {
		result = append(result, *toEventBannerDTO(&banner))
	}
	return result, nil
}
func toEventBannerDTO(banner *models.EventBanner) *dtos.EventBannerDTO {
	return &dtos.EventBannerDTO{
		ID:          banner.ID,
		Title:       banner.Title,
		Description: banner.Description,
		ImageURL:    banner.ImageURL,
		LinkURL:     banner.LinkURL,
		StartTime:   banner.StartTime,
		EndTime:     banner.EndTime,
		Priority:    banner.Priority,
		IsActive:    banner.IsActive,
		EventType:   banner.EventType,
		Position:    banner.Position,
		CreatedAt:   banner.CreatedAt,
		UpdatedAt:   banner.UpdatedAt,
	}
}
