package repositories

import (
	"github.com/TranXuanPhong25/ecom/carts/models"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

type ICartRepository interface {
	GetCart(userID string) ([]models.CartItem, error)
	GetItemQuantity(userID string, productVariantID string, shopID string) (int, error)
	AddItemToCart(item models.CartItem) error
	UpdateItemQuantity(item models.CartItem) error
	DeleteItemInCart(userID string, uuids []string) error
	ClearCart(userID string) error
}
type scyllaRepository struct {
	session *gocqlx.Session
}

func NewCartRepository() ICartRepository {
	return &scyllaRepository{
		session: &session,
	}
}
func (r *scyllaRepository) GetCart(userID string) ([]models.CartItem, error) {
	var items []models.CartItem
	query := "SELECT shop_id, product_variant_id, quantity FROM carts_ks.cart_items WHERE user_id = ?"
	q := r.session.Query(query, []string{":user_id"}).Bind(userID)
	if err := q.SelectRelease(&items); err != nil {
		return items, err
	}
	return items, nil
}

func (r *scyllaRepository) GetItemQuantity(userID string, productVariantID string, shopID string) (int, error) {
	q := r.session.Query(
		`SELECT quantity FROM carts_ks.cart_items WHERE user_id = ? AND product_variant_id = ? AND shop_id = ?`,
		[]string{":user_id", ":product_variant_id", ":shop_id"}).
		Bind(userID, productVariantID, shopID)

	quantity := 0
	if err := q.Scan(&quantity); err != nil {
		return 0, err
	}
	return quantity, nil
}

func (r *scyllaRepository) AddItemToCart(item models.CartItem) error {
	q := r.session.Query(
		`INSERT INTO carts_ks.cart_items (shop_id,product_variant_id,user_id,quantity) VALUES (?,?,?,?)`,
		[]string{":shop_id", ":product_variant_id", ":user_id", ":quantity"}).
		Bind(item.ShopID, item.ProductVariantID, item.UserID, item.Quantity)

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *scyllaRepository) UpdateItemQuantity(item models.CartItem) error {
	q := r.session.Query(
		`UPDATE carts_ks.cart_items 
				SET quantity = ? 
				WHERE user_id = ? 
  				AND product_variant_id = ?
				AND shop_id = ?`, nil).
		Bind(item.Quantity, item.UserID, item.ProductVariantID, item.ShopID)
	err := q.ExecRelease()
	if err != nil {
		return err
	}
	return nil
}

func (r *scyllaRepository) DeleteItemInCart(userID string, uuids []string) error {
	batch := r.session.Batch(gocql.LoggedBatch)
	for _, u := range uuids {
		batch.Query(
			"DELETE FROM carts_ks.cart_items WHERE user_id = ? AND product_variant_id = ?",
			userID, u,
		)
	}
	if err := r.session.ExecuteBatch(batch); err != nil {
		return err
	}
	return nil
}

func (r *scyllaRepository) ClearCart(userID string) error {
	q := r.session.Query(
		`DELETE FROM carts_ks.cart_items WHERE user_id = ?`,
		[]string{":user_id"}).Bind(userID)

	err := q.Exec()
	if err != nil {
		return err
	}
	return nil
}
