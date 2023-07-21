package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raihaninfo/books_api/initialize"
	"github.com/raihaninfo/books_api/model"
	"gorm.io/gorm"
)

var Db *gorm.DB = initialize.DbConnection()

func PostBook(c *gin.Context) {
	var book model.Books
	id := uuid.New()
	err := c.BindJSON(&book)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	Db.Create(&model.Books{Id: id, Name: book.Name, Author: book.Author, Publisher: book.Publisher, Price: book.Price})
	c.JSON(200, &book)
}

func AllBooks(c *gin.Context) {
	var books []model.Books
	Db.Find(&books)
	c.IndentedJSON(200, &books)
}

func OneBook(c *gin.Context) {
	var book model.Books
	id := c.Param("id")
	data := Db.Where("id =?", id).First(&book)
	if data.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.IndentedJSON(200, &book)
}

func UpdateBook(c *gin.Context) {
	var book model.Books
	id := c.Param("id")
	err := c.BindJSON(&book)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	data := Db.Model(&model.Books{}).Where("id =?", id).Updates(book)
	if data.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.IndentedJSON(200, &book)

}

func DeleteBook(c *gin.Context) {
	var book model.Books
	id := c.Param("id")
	data := Db.Where("id =?", id).Delete(&book)
	if data.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.IndentedJSON(200, gin.H{
		"message": "delete book " + id,
	})

}
