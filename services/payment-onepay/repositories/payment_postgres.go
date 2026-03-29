package repositories

import (
	"fmt"

	"github.com/TranXuanPhong25/ecom/services/payment-onepay/configs"
	"github.com/TranXuanPhong25/ecom/services/payment-onepay/models"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type postgresRepository struct {
	DB *gorm.DB
}

func (r *postgresRepository) CreatePayment(payment *models.Payment) error {
	if err := r.DB.Create(payment).Error; err != nil {
		log.Error("Failed to create payment:", err)
		return err
	}
	return nil
}

func (r *postgresRepository) GetByTxnRef(txnRef string) (*models.Payment, error) {
	var payment models.Payment
	if err := r.DB.Where("txn_ref = ?", txnRef).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *postgresRepository) UpdatePaymentStatus(txnRef, status, responseCode, transactionNo, message, cardType, payChannel string) error {
	tx := r.DB.Model(&models.Payment{}).
		Where("txn_ref = ?", txnRef).
		Updates(map[string]interface{}{
			"status":         status,
			"response_code":  responseCode,
			"transaction_no": transactionNo,
			"message":        message,
			"card_type":      cardType,
			"pay_channel":    payChannel,
		})
	if tx.Error != nil {
		log.Error("Failed to update payment status:", tx.Error)
		return tx.Error
	}
	return nil
}

func ConnectPostgresDB() {
	dbHost := configs.AppConfig.DBHost
	dbUser := configs.AppConfig.DBUser
	dbPassword := configs.AppConfig.DBPassword
	dbName := configs.AppConfig.DBName
	dbPort := configs.AppConfig.DBPort
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}
