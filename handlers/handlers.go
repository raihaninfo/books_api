package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raihaninfo/books_api/initialize"
	"github.com/raihaninfo/books_api/model"
	"github.com/raihaninfo/books_api/utils"
	"gorm.io/gorm"
)

var Db *gorm.DB = initialize.DbConnection()

func Login(c *gin.Context) {
	var user model.User
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Db.Where("email =?", body.Email).First(&user)
	if user.Id == uuid.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else {
		if utils.CheckPasswordHash(body.Password, user.Password) {
			tokenString := utils.CreateToken(user.Id.String())
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenString, 3600*24*2, "", "", false, true)
			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	}

}

func Logout(c *gin.Context) {

}

func Register(c *gin.Context) {
	var user model.User
	id := uuid.New()
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = id
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Db.Create(&model.User{Id: id, Name: user.Name, Email: user.Email, Username: user.Username, Password: hash})
	c.JSON(http.StatusCreated, &user)

}

func PostBook(c *gin.Context) {
	var book model.Books
	id := uuid.New()
	err := c.BindJSON(&book)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	book.Id = id
	Db.Create(&model.Books{Id: id, Name: book.Name, Author: book.Author, Publisher: book.Publisher, Price: book.Price})
	c.JSON(201, &book)
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
