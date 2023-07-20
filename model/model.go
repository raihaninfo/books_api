package model

import (
	"github.com/google/uuid"
	"github.com/raihaninfo/books_api/initialize"

	"gorm.io/gorm"
)

type Books struct {
	Id        uuid.UUID
	Name      string
	Author    string
	Publisher string
	Price     float32
}

// db is a gorm database connection
var Db *gorm.DB = initialize.DbConnection()

func InsertSingleBook(name, author, publisher string, price float32) {
	id := uuid.New()
	Db.Create(&Books{Id: id, Name: name, Author: author, Publisher: publisher, Price: price})
}
