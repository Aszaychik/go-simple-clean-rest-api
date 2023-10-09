package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() (*gorm.DB, error) {
	// godotenv.Load() 
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load environment variables: %w", err)
	// }

	return initDB()
}

func initDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}