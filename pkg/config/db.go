package config

import (
	// "os"

	"github.com/Bek-arys/BookStore/pkg/models"
	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// errr := godotenv.Load("C:/Users/HP/Desktop/assignment-3/bookstore/.env")

	// if errr != nil {
	// 	panic(errr)
	// }
	
	db, err := gorm.Open(postgres.Open("postgresql://postgres:Bek@rys_2003@postgres:5432/assignment_3"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})
	DB = db
}