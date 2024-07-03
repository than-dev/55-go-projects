package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var (
	db * gorm.DB
)

func ConnectDB() {
	d, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})


	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}