package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	TxnRef        string    `gorm:"type:varchar(40);not null;uniqueIndex"` // vpc_MerchTxnRef
	OrderID       string    `gorm:"type:varchar(40);not null"`
	Amount        int64     `gorm:"type:bigint;not null"` // in VND (original, before *100)
	Status        string    `gorm:"type:varchar(20);not null;default:'pending'"` // pending, success, failed
	ResponseCode  string    `gorm:"type:varchar(10)"`
	TransactionNo string    `gorm:"type:varchar(32)"`
	Message       string    `gorm:"type:varchar(200)"`
	CardType      string    `gorm:"type:varchar(20)"`
	PayChannel    string    `gorm:"type:varchar(20)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
