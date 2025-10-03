package services

import (
	"github.com/TranXuanPhong25/ecom/carts/dtos"
	"github.com/TranXuanPhong25/ecom/carts/models"
	"github.com/TranXuanPhong25/ecom/carts/repositories"
)

type ICartService interface {
	GetCart(userID string) (*dtos.Cart, error)
	AddItemToCart(userID string, item *dtos.CartItemPayload) error
	UpdateItemInCart(userID string, item *dtos.CartItemPayload) error
	DeleteItemInCart(userID string, item *dtos.CartItemPayload) error
	ClearCart(userID string) error
}
type CartService struct {
	repo           repositories.ICartRepository
	productService IProductService
}

func NewCartService(repo repositories.ICartRepository, productService IProductService) ICartService {

	return &CartService{
		repo:           repo,
		productService: productService,
	}
}

func (s *CartService) GetCart(userID string) (*dtos.Cart, error) {

	items, err := s.repo.GetCart(userID)
	if err != nil {
		return nil, err
	}
	productVariantIDs := make([]string, len(items))
	for i, item := range items {
		productVariantIDs[i] = item.ProductVariantID
	}
	productVariantsResponse, err := s.productService.GetProductVariantByIds(productVariantIDs)
	if err != nil {
		return nil, err
	}
	//TODO: handle not found product variants
	// if len(productVariantsResponse.NotFoundIDs) > 0 {
	// 	for _, notFoundID := range productVariantsResponse.NotFoundIDs {
	// 		for _, item := range items {
	// 			if item.ProductVariantID == notFoundID {
	// 				_ = s.repo.DeleteItemInCart(item)
	// 			}
	// 		}
	// 	}
	// }
	// remove notfound in items
	cartItems := make([]dtos.CartItem, 0, len(productVariantsResponse.Variants))
	for i, item := range items {
		cartItems = append(cartItems, dtos.CartItem{
			ProductVariant: productVariantsResponse.Variants[i],
			Quantity:       item.Quantity,
			ShopID:         item.ShopID,
		})
	}
	cart := dtos.Cart{
		Items: cartItems,
	}
	return &cart, nil
}

func (s *CartService) AddItemToCart(userID string, item *dtos.CartItemPayload) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	existingQuantity, err := s.repo.GetItemQuantity(userID, item.ProductVariantID, item.ShopID)
	if (existingQuantity > 0) || (err == nil) {
		cartItem.Quantity += existingQuantity
		return s.repo.UpdateItemQuantity(cartItem)
	}
	return s.repo.AddItemToCart(cartItem)
}

func (s *CartService) UpdateItemInCart(userID string, item *dtos.CartItemPayload) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	return s.repo.UpdateItemQuantity(cartItem)
}

func (s *CartService) DeleteItemInCart(userID string, item *dtos.CartItemPayload) error {
	cartItem := models.CartItem{
		UserID:           userID,
		ProductVariantID: item.ProductVariantID,
		ShopID:           item.ShopID,
		Quantity:         item.Quantity,
	}
	return s.repo.DeleteItemInCart(cartItem)
}

func (s *CartService) ClearCart(userID string) error {
	return s.repo.ClearCart(userID)
}
