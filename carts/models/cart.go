package models

// CREATE KEYSPACE IF NOT EXISTS ecommerce WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' };
type CartItem struct {
	UserID    string
	CartID    string
	ProductID string
	ShopID    string
	Quantity  int
}

//cdc = {'enabled': true, 'preimage': true, 'postimage': true};

type Cart struct {
	UserID   string
	CartId   string
	IsActive bool
}

//CREATE INDEX cart_is_active ON ecommerce.cart ((user_id), is_active);
