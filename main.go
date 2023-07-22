package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/raihaninfo/books_api/initialize"
	"github.com/raihaninfo/books_api/model"
	"github.com/raihaninfo/books_api/router"
)

func init() {
	initialize.Env()
	db := initialize.DbConnection()
	db.AutoMigrate(&model.Books{}, &model.User{})
}
func main() {
	r := gin.Default()

	router.Router(r)

	r.Run()
}
