package repositories

import (
	"strings"

	"github.com/TranXuanPhong25/ecom/carts/configs"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var (
	session gocqlx.Session
)

func InitDBConnection() {
	ConnectScyllaDB()
	MigrateDB()
}
func ConnectScyllaDB() {
	rawScyllaNodes := configs.AppConfig.ScyllaNodes
	scyllaNodes := strings.Split(rawScyllaNodes, ",")

	cluster := gocql.NewCluster(scyllaNodes...)
	cluster.Port = 9042
	//cluster.Keyspace = "carts_ks"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "password123"}
	//cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")
	newSession, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic("Connection to scyllaDB failed")
	}
	session = newSession
}

func MigrateDB() {
	session.Query("DROP KEYSPACE IF EXISTS carts_ks", nil).Exec()
	session.Query(
		"CREATE KEYSPACE IF NOT EXISTS carts_ks "+
			"WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': '3'}  AND durable_writes = true AND TABLETS = {'enabled': false};;",
		nil).Exec()
	session.Query(
		"CREATE TABLE IF NOT EXISTS "+
			"carts_ks.carts( "+
			"user_id uuid,"+
			"cart_id uuid,"+
			"shop_id uuid,"+
			"product_variant_id uuid,"+
			"quantity int,"+
			"PRIMARY KEY (cart_id, user_id, product_variant_id))",
		nil).Exec()
	session.Query(
		"CREATE INDEX IF NOT EXISTS cart_is_active "+
			"ON carts_ks.carts (is_active)",
		nil).Exec()
	session.Query(
		"CREATE INDEX IF NOT EXISTS cart_product_variant_id_idx "+
			"ON carts_ks.carts (product_variant_id)",
		nil).Exec()
	session.Query("CREATE TABLE IF NOT EXISTS "+
		"carts_ks.cart_items( "+
		"user_id uuid,"+
		"cart_id uuid,"+
		"is_active text,"+
		"PRIMARY KEY (user_id, cart_id))"+
		"WITH WITH cdc = {'enabled': true, 'preimage': true, 'postimage': true}",
		nil).Exec()
}
