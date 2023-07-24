package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/books_api/handlers"
	"github.com/raihaninfo/books_api/middleware"
)

func Router(r *gin.Engine) {

	public := r.Group("/api")
	public.POST("/register", handlers.Register)
	public.POST("/login", handlers.Login)

	api := r.Group("/api/")
	api.Use(middleware.RequirerAuth)
	api.POST("/book", handlers.PostBook)
	api.GET("/books", handlers.AllBooks)
	api.GET("/book/:id", handlers.OneBook)
	api.PUT("/book/:id", handlers.UpdateBook)
	api.DELETE("/book/:id", handlers.DeleteBook)
}
