package models

import (
	"github.com/google/uuid"
)

// CREATE KEYSPACE IF NOT EXISTS ecommerce WITH replication = { 'class': 'NetworkTopologyStrategy', 'replication_factor': '3' };
type CartItem struct {
	UserID           uuid.UUID `gorm:"type:uuid;not null;index:idx_cart_item,unique"`
	ProductVariantID int       `gorm:"type:int;not null;index:idx_cart_item,unique"`
	ShopID           uuid.UUID `gorm:"type:uuid;not null;index:idx_cart_item,unique"`
	Quantity         int       `gorm:"type:int;not null"`
}

//cdc = {'enabled': true, 'preimage': true, 'postimage': true};
