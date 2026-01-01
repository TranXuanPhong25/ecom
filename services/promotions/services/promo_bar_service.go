package services

import (
	"net/http"
	"time"

	"github.com/TranXuanPhong25/ecom/services/promotions/dtos"
	"github.com/TranXuanPhong25/ecom/services/promotions/models"
	"github.com/TranXuanPhong25/ecom/services/promotions/repositories"
	"github.com/labstack/echo/v4"
)

// ============ PROMO BAR SERVICES ============

func CreatePromoBar(request *models.CreatePromoBarRequest) (*dtos.PromoBarDTO, *echo.HTTPError) {
	promoBar := &models.PromoBar{
		Message:         request.Message,
		BackgroundColor: request.BackgroundColor,
		TextColor:       request.TextColor,
		LinkURL:         request.LinkURL,
		StartTime:       request.StartTime,
		EndTime:         request.EndTime,
		IsActive:        request.IsActive,
		Priority:        request.Priority,
		IsCloseable:     request.IsCloseable,
	}

	if promoBar.BackgroundColor == "" {
		promoBar.BackgroundColor = "#ff0000"
	}
	if promoBar.TextColor == "" {
		promoBar.TextColor = "#ffffff"
	}

	tx := repositories.DB.Create(promoBar)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while creating promo bar")
	}

	return toPromoBarDTO(promoBar), nil
}

func GetPromoBarByID(id string) (*dtos.PromoBarDTO, *echo.HTTPError) {
	promoBar := &models.PromoBar{}
	tx := repositories.DB.First(promoBar, "id = ?", id)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Promo bar not found")
	}
	return toPromoBarDTO(promoBar), nil
}

func GetAllPromoBars(activeOnly bool) ([]dtos.PromoBarDTO, *echo.HTTPError) {
	var promoBars []models.PromoBar
	query := repositories.DB.Order("priority DESC, created_at DESC")

	if activeOnly {
		now := time.Now()
		query = query.Where("is_active = ? AND start_time <= ? AND end_time >= ?", true, now, now)
	}

	tx := query.Find(&promoBars)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while fetching promo bars")
	}

	result := make([]dtos.PromoBarDTO, 0)
	for _, bar := range promoBars {
		result = append(result, *toPromoBarDTO(&bar))
	}
	return result, nil
}

func UpdatePromoBar(request *models.UpdatePromoBarRequest) (*dtos.PromoBarDTO, *echo.HTTPError) {
	promoBar := &models.PromoBar{}
	tx := repositories.DB.First(promoBar, "id = ?", request.ID)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Promo bar not found")
	}

	updates := make(map[string]interface{})
	if request.Message != "" {
		updates["message"] = request.Message
	}
	if request.BackgroundColor != "" {
		updates["background_color"] = request.BackgroundColor
	}
	if request.TextColor != "" {
		updates["text_color"] = request.TextColor
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
	if request.IsActive != nil {
		updates["is_active"] = *request.IsActive
	}
	if request.Priority != nil {
		updates["priority"] = *request.Priority
	}
	if request.IsCloseable != nil {
		updates["is_closeable"] = *request.IsCloseable
	}

	tx = repositories.DB.Model(promoBar).Updates(updates)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while updating promo bar")
	}

	return toPromoBarDTO(promoBar), nil
}

func DeletePromoBar(id string) *echo.HTTPError {
	tx := repositories.DB.Delete(&models.PromoBar{}, "id = ?", id)
	if tx.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error while deleting promo bar")
	}
	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Promo bar not found")
	}
	return nil
}

// ============ COMBINED SERVICES ============

func GetActivePromotions() (*dtos.GetActivePromotionsResponse, *echo.HTTPError) {
	banners, err1 := GetAllEventBanners(true)
	if err1 != nil {
		return nil, err1
	}

	promoBars, err2 := GetAllPromoBars(true)
	if err2 != nil {
		return nil, err2
	}

	return &dtos.GetActivePromotionsResponse{
		EventBanners: banners,
		PromoBars:    promoBars,
	}, nil
}

// ============ HELPER FUNCTIONS ============

func toPromoBarDTO(promoBar *models.PromoBar) *dtos.PromoBarDTO {
	return &dtos.PromoBarDTO{
		ID:              promoBar.ID,
		Message:         promoBar.Message,
		BackgroundColor: promoBar.BackgroundColor,
		TextColor:       promoBar.TextColor,
		LinkURL:         promoBar.LinkURL,
		StartTime:       promoBar.StartTime,
		EndTime:         promoBar.EndTime,
		IsActive:        promoBar.IsActive,
		Priority:        promoBar.Priority,
		IsCloseable:     promoBar.IsCloseable,
		CreatedAt:       promoBar.CreatedAt,
		UpdatedAt:       promoBar.UpdatedAt,
	}
}
