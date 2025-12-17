package models

// CREATE KEYSPACE IF NOT EXISTS ecommerce WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' };
type CartItem struct {
	UserID           string
	ProductVariantID string
	ShopID           string
	Quantity         int
}

//cdc = {'enabled': true, 'preimage': true, 'postimage': true};
