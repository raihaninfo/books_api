package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/books_api/handlers"
)

func Router(r *gin.Engine) {

	public := r.Group("/api")
	public.POST("/register", handlers.Register)
	public.POST("/login", handlers.Login)
	public.GET("/logout", handlers.Logout)

	api := r.Group("/api/")
	api.POST("/book", handlers.PostBook)
	api.GET("/books", handlers.AllBooks)
	api.GET("/book/:id", handlers.OneBook)
	api.PUT("/book/:id", handlers.UpdateBook)
	api.DELETE("/book/:id", handlers.DeleteBook)
}
