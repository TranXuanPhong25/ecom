package repositories

import "github.com/TranXuanPhong25/ecom/carts/models"

func GetCart(userID string) ([]models.CartItem, error) {
	var items []models.CartItem
	names := []string{userID}
	query := "SELECT shop_id, product_id, quantity FROM cart_items WHERE user_id = ?"
	iter := session.Query(query, names).Iter()
	var c models.CartItem
	for iter.Scan(&c.ShopID, &c.ProductID, &c.Quantity) {
		items = append(items, c)
	}
	return items, nil
}
