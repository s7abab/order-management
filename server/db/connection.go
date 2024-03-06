package db

import (
	"orders/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDb() (*gorm.DB, error) {
	dsn := "root:2392@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	DB = db
	// Auto migrate tables
	err = db.AutoMigrate(
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDb() *gorm.DB {
	return DB
}
