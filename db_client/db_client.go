package db_client

import (
	"fmt"
	"payment-service/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDBConnection() {

	db, err := gorm.Open(postgres.Open("postgres://postgres:secret@postgres:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.Payment{})
	DB = db

	fmt.Printf("\nSuccessfully connected to database!\n")
}
