package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	const connStr = "postgres://postgres:1234@localhost:5432/go-restapi"
	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	db.AutoMigrate(&User{})

	DB = db
}
