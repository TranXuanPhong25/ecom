package repositories

import (
	"fmt"
	"log"

	"github.com/TranXuanPhong25/ecom/services/shops/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	dbHost := configs.AppConfig.DBHost
	dbUser := configs.AppConfig.DBUser
	dbPassword := configs.AppConfig.DBPassword
	dbName := configs.AppConfig.DBName
	dbPort := configs.AppConfig.DBPort
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}
