package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/books_api/model"
)

func PostBook(c *gin.Context) {
	var book model.Books
	err := c.BindJSON(book)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	model.InsertSingleBook(book.Name, book.Author, book.Publisher, book.Price)
}

func GetAllBooks(c *gin.Context) {

}

func GetOneBook(c *gin.Context) {

}

func UpdateBook(c *gin.Context) {

}
func DeleteBook(c *gin.Context) {

}
