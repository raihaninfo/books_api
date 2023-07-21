package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/books_api/handlers"
)

func Router(r *gin.Engine) {
	r.POST("/book", handlers.PostBook)
	r.GET("/books", handlers.AllBooks)
	r.GET("/book/:id", handlers.OneBook)
	r.PUT("/book/:id", handlers.UpdateBook)
	r.DELETE("/book/:id", handlers.DeleteBook)
}

