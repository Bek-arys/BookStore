package controllers

import (
	"fmt"
	"strings"

	"github.com/Bek-arys/assignment-3/pkg/config"
	"github.com/Bek-arys/assignment-3/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := []models.Book{}
	config.DB.Find(&books)
	c.JSON(200, &books)
}

func GetBook(c *gin.Context) {
	var book models.Book
	config.DB.Where("id = ?", c.Param("id")).First(&book)
	if (book == models.Book{}) {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	}
	c.JSON(200, &book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	if book.Title == "" || book.Cost == 0 {
		c.JSON(400, gin.H{
			"detail": "Post the 'title' and 'cost' fields in the request body.",
		})
		return
	}
	config.DB.Create(&book)
	c.JSON(200, &book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	config.DB.Where("id = ?", c.Param("id")).Delete(&book)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	config.DB.First(&book, id)
	if (book == models.Book{}) {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	}
	c.BindJSON(&book)
	config.DB.Save(&book)
}

func SearchBookByTitle(c *gin.Context) {
	var title string = c.Query("q")
	books := []models.Book{}
	config.DB.Where("LOWER(title) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(title))).Find(&books)
	c.JSON(200, &books)
}

func SortedBooks(c *gin.Context) {
	var order string = c.Query("q")
	order = strings.ToLower(order)
	books := []models.Book{}
	if !(order == "desc" || order == "asc") {
		c.JSON(400, gin.H{
			"message": "Invalid Query format",
		})
		return
	}

	config.DB.Order(fmt.Sprintf("cost %s", order)).Find(&books)
	c.JSON(200, &books)
}
