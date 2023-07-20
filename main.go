package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	db "github.com/raihaninfo/books_api/initialize"
	"github.com/raihaninfo/books_api/model"
)

func main() {
	db := db.DbConnection()
	db.AutoMigrate(&model.Books{})
	r := gin.Default()


	r.Run(":8080")
}
