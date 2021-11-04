package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	// TODO: USE ENV VAR
	dsn := "host=localhost user=postgres password=postgres dbname=oneline-shop port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
