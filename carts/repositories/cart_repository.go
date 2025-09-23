package repositories

import (
	"github.com/TranXuanPhong25/ecom/carts/models"
)

func GetCart(userID string) ([]models.CartItem, error) {
	var items []models.CartItem
	query := "SELECT shop_id, product_variant_id, quantity FROM carts_ks.cart_items WHERE user_id = ?"
	iter := session.Query(query, []string{":user_id"}).
		BindMap(map[string]interface{}{
			":user_id": userID,
		}).Iter()
	var c models.CartItem
	for iter.Scan(&c.ShopID, &c.ProductVariantID, &c.Quantity) {
		items = append(items, c)
	}
	return items, nil
}

func GetItemQuantity(userID string, productVariantID string) (int, error) {
	q := session.Query(
		`SELECT quantity FROM carts_ks.cart_items WHERE product_variant_id = ? AND user_id = ?`,
		[]string{":user_id", ":product_variant_id"}).
		BindMap(map[string]interface{}{
			":product_variant_id": productVariantID,
			":user_id":            userID,
		})
	var existingItem models.CartItem
	if err := q.SelectRelease(&existingItem); err != nil {
		return 0, err
	}
	return existingItem.Quantity, nil
}

func AddItemToCart(item models.CartItem) error {
	q := session.Query(
		`INSERT INTO carts_ks.cart_items (shop_id,product_variant_id,user_id,quantity) VALUES (?,?,?,?)`,
		[]string{":shop_id", ":product_variant_id", ":user_id", ":quantity"}).
		BindMap(map[string]interface{}{
			":shop_id":            item.ShopID,
			":product_variant_id": item.ProductVariantID,
			":user_id":            item.UserID,
			":quantity":           item.Quantity,
		})

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}

func UpdateItemQuantity(item models.CartItem) error {
	q := session.Query(
		`UPDATE carts_ks.cart_items SET quantity = ? WHERE product_variant_id = ? AND user_id = ?`,
		[]string{":quantity", ":product_variant_id", ":user_id"}).
		BindMap(map[string]interface{}{
			":quantity":           item.Quantity,
			":product_variant_id": item.ProductVariantID,
			":user_id":            item.UserID,
		})

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}

func DeleteItemInCart(item models.CartItem) error {
	q := session.Query(
		`DELETE FROM carts_ks.cart_items WHERE product_variant_id = ? AND user_id = ?`,
		[]string{":product_variant_id", ":user_id"}).
		BindMap(map[string]interface{}{
			":product_variant_id": item.ProductVariantID,
			":user_id":            item.UserID,
		})

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
		BindMap(map[string]interface{}{
			":user_id": userID,
		})

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}
