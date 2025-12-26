package services

import (
	"github.com/TranXuanPhong25/ecom/services/carts/dtos"
	"github.com/TranXuanPhong25/ecom/services/carts/models"
	"github.com/TranXuanPhong25/ecom/services/carts/repositories"
	"github.com/google/uuid"
)

type ICartService interface {
	GetCart(userID string) (*dtos.Cart, error)
	AddItemToCart(userID uuid.UUID, item *dtos.CartItemPayload) error
	UpdateItemInCart(userID uuid.UUID, item *dtos.CartItemPayload) error
	DeleteItemInCart(userID string, itemIds []int) error
	ClearCart(userID string) error
	GetTotalItems(id string) (int, error)
}
type CartService struct {
	repo           repositories.ICartRepository
	productService IProductService
	shopsService   IShopsService
}

func NewCartService(
	repo repositories.ICartRepository,
	productService IProductService,
	shopsService IShopsService,
) ICartService {

	return &CartService{
		repo:           repo,
		productService: productService,
		shopsService:   shopsService,
	}
}

func (s *CartService) GetCart(userID string) (*dtos.Cart, error) {

	items, err := s.repo.GetCart(userID)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return &dtos.Cart{
			Items: []dtos.CartItem{},
		}, nil
	}
	cartItems, err := s.getProductVariantsFromCartItems(userID, &items)
	if err != nil {
		return nil, err
	}

	shops, err := s.getUniqueShopsFromCartItems(&items)
	if err != nil {
		return nil, err
	}
	cart := dtos.Cart{
		Items: cartItems,
		Shops: shops,
	}
	return &cart, nil
}

func (s *CartService) getUniqueShopsFromCartItems(items *[]models.CartItem) ([]dtos.Shop, error) {
	seen := make(map[string]struct{})
	var uniqueShopIDs []string
	for _, item := range *items {
		if _, ok := seen[item.ShopID.String()]; !ok {
			seen[item.ShopID.String()] = struct{}{}
			uniqueShopIDs = append(uniqueShopIDs, item.ShopID.String())
		}
	}

	if len(uniqueShopIDs) == 0 {
		return []dtos.Shop{}, nil
	}
	shopsResponse, err := s.shopsService.GetShopsByIds(uniqueShopIDs)
	if err != nil {
		return nil, err
	}
	return shopsResponse.Shops, nil
}

func (s *CartService) getProductVariantsFromCartItems(userID string, items *[]models.CartItem) ([]dtos.CartItem, error) {
	productVariantIDs := make([]int, len(*items))
	for i, item := range *items {
		productVariantIDs[i] = item.ProductVariantID
	}
	getProductVariantsResponse, err := s.productService.GetProductVariantByIds(productVariantIDs)
	if err != nil {
		return nil, err
	}
	if len(getProductVariantsResponse.NotFoundIDs) > 0 {
		//remove not found items from cart
		err := s.repo.DeleteItemInCart(userID, getProductVariantsResponse.NotFoundIDs)
		if err != nil {
			return nil, err
		}
	}

	productVariants := getProductVariantsResponse.Variants
	//build map of productVariantID to ProductVariant
	productVariantMap := make(map[int]dtos.ProductVariant)
	for _, pv := range productVariants {
		productVariantMap[pv.ID] = pv
	}
	cartItems := make([]dtos.CartItem, len(productVariants))
	for i, item := range *items {
		cartItems[i] = dtos.CartItem{
			ProductVariant: productVariantMap[item.ProductVariantID],
			Quantity:       item.Quantity,
			ShopID:         item.ShopID.String(),
		}
	}
	return cartItems, nil
}
func (s *CartService) AddItemToCart(userID uuid.UUID, item *dtos.CartItemPayload) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	existingQuantity, err := s.repo.GetItemQuantity(userID.String(), item.ProductVariantID, item.ShopID.String())
	if (existingQuantity > 0) || (err == nil) {
		cartItem.Quantity += existingQuantity
		return s.repo.UpdateItemQuantity(cartItem)
	}
	return s.repo.AddItemToCart(cartItem)
}

func (s *CartService) UpdateItemInCart(userID uuid.UUID, item *dtos.CartItemPayload) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	return s.repo.UpdateItemQuantity(cartItem)
}

func (s *CartService) DeleteItemInCart(userID string, itemIds []int) error {
	return s.repo.DeleteItemInCart(userID, itemIds)
}

func (s *CartService) ClearCart(userID string) error {
	return s.repo.ClearCart(userID)
}

func (s *CartService) GetTotalItems(id string) (int, error) {
	return s.repo.GetTotalItemsInCart(id)
}
