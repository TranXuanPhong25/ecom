package repositories

import (
	"github.com/TranXuanPhong25/ecom/carts/models"
)

func GetCart(userID string) ([]models.CartItem, error) {
	var items []models.CartItem
	query := "SELECT shop_id, product_variant_id, quantity FROM carts_ks.cart_items WHERE user_id = ?"
	q := session.Query(query, []string{":user_id"}).
		Bind(userID)
	if err := q.SelectRelease(&items); err != nil {
		return items, err
	}
	return items, nil
}

func GetItemQuantity(userID string, productVariantID string) (int, error) {
	q := session.Query(
		`SELECT quantity FROM carts_ks.cart_items WHERE user_id = ? AND product_variant_id = ? `,
		[]string{":user_id", ":product_variant_id"}).
		Bind(userID, productVariantID)

	var existingItem models.CartItem
	if err := q.GetRelease(&existingItem); err != nil {
		return 0, err
	}
	return existingItem.Quantity, nil
}

func AddItemToCart(item models.CartItem) error {
	q := session.Query(
		`INSERT INTO carts_ks.cart_items (shop_id,product_variant_id,user_id,quantity) VALUES (?,?,?,?)`,
		[]string{":shop_id", ":product_variant_id", ":user_id", ":quantity"}).
		Bind(item.ShopID, item.ProductVariantID, item.UserID, item.Quantity)

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}

func UpdateItemQuantity(item models.CartItem) error {
	q := session.Query(
		`UPDATE carts_ks.cart_items SET quantity = ? WHERE user_id = ? AND product_variant_id = ?`,
		[]string{":quantity", ":user_id", ":product_variant_id"}).
		Bind(item.Quantity, item.UserID, item.ProductVariantID)

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteItemInCart(item models.CartItem) error {
	q := session.Query(
		`DELETE FROM carts_ks.cart_items WHERE user_id = ? AND product_variant_id = ?`,
		[]string{":user_id", ":product_variant_id"}).
		Bind(item.UserID, item.ProductVariantID)

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}

func ClearCart(userID string) error {
	q := session.Query(
		`DELETE FROM carts_ks.cart_items WHERE user_id = ?`,
		[]string{":user_id"}).
		Bind(userID)

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}
