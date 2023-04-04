package routes

import (
	"github.com/Bek-arys/assignment-3/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoute(r *gin.Engine) {
	r.GET("books/", controllers.GetBooks)
	r.GET("books/:id", controllers.GetBook)
	r.GET("books/search", controllers.SearchBookByTitle)
	r.GET("books/sort", controllers.SortedBooks)
	r.POST("books/", controllers.CreateBook)
	r.PUT("books/:id", controllers.UpdateBook)
	r.DELETE("books/:id", controllers.DeleteBook)
}