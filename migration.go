package main

import (
	"github.com/negadive/oneline/db"
	"github.com/negadive/oneline/model"
)

func Migrate() {
	_db := db.GetDb()

	_db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderProduct{})
}
