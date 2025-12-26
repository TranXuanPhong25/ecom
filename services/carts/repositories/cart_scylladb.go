package repositories

import (
	"fmt"
	"strings"

	"github.com/TranXuanPhong25/ecom/services/carts/configs"
	"github.com/TranXuanPhong25/ecom/services/carts/models"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
	"github.com/scylladb/gocqlx/v3/table"
)

var (
	session gocqlx.Session
	stmts   = createStatements()
)

func InitScyllaDBConnection() {
	ConnectScyllaDB()
	MigrateDB()
}
func ConnectScyllaDB() {
	rawScyllaNodes := configs.AppConfig.ScyllaNodes
	scyllaNodes := strings.Split(rawScyllaNodes, ",")

	cluster := gocql.NewCluster(scyllaNodes...)
	cluster.Port = 9042
	//cluster.Keyspace = "carts_ks"
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "password123"}
	//cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")
	newSession, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Printf("Error connecting to ScyllaDB: %v\n", err)
		panic("Connection to scyllaDB failed")
	}
	session = newSession
}

func MigrateDB() {
	//err := session.Query("DROP KEYSPACE IF EXISTS carts_ks", nil).Exec()
	//if err != nil {
	//	panic(err)
	//}
	err := session.Query(
		"CREATE KEYSPACE IF NOT EXISTS carts_ks "+
			"WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': '1'}  AND durable_writes = true AND TABLETS = {'enabled': false};;",
		nil).Exec()
	if err != nil {
		panic(err)
	}
	err = session.Query(
		"CREATE TABLE IF NOT EXISTS "+
			"carts_ks.cart_items( "+
			"user_id uuid,"+
			"shop_id uuid,"+
			"product_variant_id int,"+
			"quantity int,"+
			"PRIMARY KEY ((user_id), product_variant_id, shop_id))"+
			"WITH cdc = {'enabled': true, 'preimage': true, 'postimage': true}",
		nil).Exec()
	if err != nil {
		panic(err)
	}

}

func createStatements() *statements {
	m := table.Metadata{
		Name:    "cart_items",
		Columns: []string{"user_id", "shop_id", "product_variant_id", "quantity"},
		PartKey: []string{"user_id"},
	}
	tbl := table.New(m)

	deleteStmt, deleteNames := tbl.Delete()
	insertStmt, insertNames := tbl.Insert()
	updateStmt, updateNames := tbl.Update()
	// Normally a select statement such as this would use `tbl.Select()` to select by
	// primary key but now we just want to display all the records...
	selectStmt, selectNames := qb.Select(m.Name).Columns(m.Columns...).ToCql()
	return &statements{
		del: query{
			stmt:  deleteStmt,
			names: deleteNames,
		},
		ins: query{
			stmt:  insertStmt,
			names: insertNames,
		},
		sel: query{
			stmt:  selectStmt,
			names: selectNames,
		},
		upd: query{
			stmt:  updateStmt,
			names: updateNames,
		},
	}
}

type query struct {
	stmt  string
	names []string
}

type statements struct {
	del query
	ins query
	upd query
	sel query
}
type scyllaRepository struct {
	session *gocqlx.Session
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
