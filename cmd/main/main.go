package main

import (
	"github.com/Bek-arys/BookStore/pkg/config"
	"github.com/Bek-arys/BookStore/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.BookRoute(router)

	router.Run(":8080")
}