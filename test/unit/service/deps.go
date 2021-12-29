package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/negadive/oneline/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=oneline-shop-test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&model.User{}, &model.Product{})

	return db
}

func truncate_user(db *gorm.DB) {
	db.Unscoped().Where("1 = 1").Delete(&model.User{})
}

func truncate_product(db *gorm.DB) {
	db.Unscoped().Where("1 = 1").Delete(&model.Product{})
}

func TestMain(m *testing.M) {

	db := setupTestDB()
	fmt.Println("TestMain")
	db.AutoMigrate(&model.User{}, &model.Product{})

	m.Run()
}
