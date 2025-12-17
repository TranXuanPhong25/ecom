package services

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/TranXuanPhong25/ecom/services/shops/dtos"
	"github.com/TranXuanPhong25/ecom/services/shops/models"
	pb "github.com/TranXuanPhong25/ecom/services/shops/proto"
	"github.com/TranXuanPhong25/ecom/services/shops/repositories"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type RPCShopsService struct {
	pb.UnimplementedShopsServiceServer
}

func (s *RPCShopsService) GetShopsByIDs(_ context.Context, in *pb.GetShopsByIDsRequest) (*pb.GetShopsByIDsResponse, error) {
	response, err := GetShopsByIDs(in.GetIds())
	if err != nil {
		return nil, err
	}
	var shopProtos []*pb.Shop
	for _, shop := range response.Shops {
		shopProtos = append(shopProtos, &pb.Shop{
			Id:   shop.ID.String(),
			Name: shop.Name,
		})
	}
	return &pb.GetShopsByIDsResponse{
		Shops:       shopProtos,
		NotFoundIds: response.NotFoundIDs,
	}, nil
}

func RegisterService(s *grpc.Server) {
	pb.RegisterShopsServiceServer(s, &RPCShopsService{})
	log.Info("ShopService registered")
}

func CreateShop(request *models.CreateShopRequest) (*dtos.ShopDTO, *echo.HTTPError) {
	shop := &models.Shop{
		Name:         request.Name,
		OwnerID:      request.OwnerID,
		Location:     request.Location,
		Logo:         request.Logo,
		Banner:       request.Banner,
		Email:        request.Email,
		Phone:        request.Phone,
		BusinessType: request.BusinessType,
	}
	tx := repositories.DB.Create(shop)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "unique constraint") {
			return &dtos.ShopDTO{}, echo.NewHTTPError(http.StatusBadRequest, "Shop with this owner already exists")
		}
		return &dtos.ShopDTO{}, echo.NewHTTPError(http.StatusInternalServerError, "Error while creating shop")
	}
	return toShopDTO(shop), nil

}

func GetShopsByOwnerID(ownerId string) (*dtos.ShopDTO, *echo.HTTPError) {
	shop := &models.Shop{}
	tx := repositories.DB.First(shop, "owner_id = ?", ownerId)

	if tx.Error != nil {
		// Kiểm tra nếu là lỗi record not found
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Shop not found")
		}
		// Các lỗi khác (database connection, SQL syntax, etc.)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, tx.Error.Error())
	}

	return toShopDTO(shop), nil
}

func GetShopsByIDs(shopIDs []string) (*dtos.GetShopsResponse, *echo.HTTPError) {
	var shops []models.Shop
	tx := repositories.DB.Where("id IN ?", shopIDs).Find(&shops)
	if tx.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while fetching shops")
	}
	foundShopIDs := make(map[string]bool)
	for _, shop := range shops {
		foundShopIDs[shop.ID.String()] = true
	}

	var notFoundIDs []string
	for _, id := range shopIDs {
		if !foundShopIDs[id] {
			notFoundIDs = append(notFoundIDs, id)
		}
	}

	var shopDTOs []dtos.ShopDTO
	for _, shop := range shops {
		shopDTOs = append(shopDTOs, *toShopDTO(&shop))
	}

	response := &dtos.GetShopsResponse{
		Shops:       shopDTOs,
		NotFoundIDs: notFoundIDs,
	}

	return response, nil
}
func toShopDTO(shop *models.Shop) *dtos.ShopDTO {
	return &dtos.ShopDTO{
		ID:           shop.ID,
		Name:         shop.Name,
		OwnerID:      shop.OwnerID,
		Location:     shop.Location,
		Rating:       shop.Rating,
		Logo:         shop.Logo,
		Banner:       shop.Banner,
		Email:        shop.Email,
		Phone:        shop.Phone,
		BusinessType: shop.BusinessType,
	}
}
