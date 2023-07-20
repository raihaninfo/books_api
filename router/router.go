package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/books_api/handlers"
)

func Router(r *gin.Engine) {
	r.POST("/book", handlers.PostBook)
	r.GET("/books", handlers.GetAllBooks)
	r.GET("/book/{:id}", handlers.GetOneBook)
	r.PUT("/book/{:id}", handlers.GetOneBook)
	r.DELETE("/book/{:id}", handlers.DeleteBook)
}
