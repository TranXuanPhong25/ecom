package repositories

import "github.com/TranXuanPhong25/ecom/services/payment-onepay/models"

type IPaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetByTxnRef(txnRef string) (*models.Payment, error)
	UpdatePaymentStatus(txnRef, status, responseCode, transactionNo, message, cardType, payChannel string) error
}

func InitDBConnection() {
	ConnectPostgresDB()
	if err := DB.AutoMigrate(&models.Payment{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

func NewPaymentRepository() IPaymentRepository {
	return &postgresRepository{DB: DB}
}
