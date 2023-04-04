package config

import (
	"os"

	"github.com/Bek-arys/BookStore/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	errr := godotenv.Load("../../.env")

	if errr != nil {
		panic(errr)
	}
	
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_SOURCE")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})
	DB = db
}