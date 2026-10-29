package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=142005 dbname=ProtectingYourLife port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("khong the ket noi DB", err)
	}

	log.Println("Ket Noi DB Thanh Cong")

	return db
}