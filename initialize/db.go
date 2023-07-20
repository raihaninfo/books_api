package initialize

import (
	"log"

	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConnection() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate(db)

	log.Println("db connection successful")
	return db
}

// func migrate(db *gorm.DB) {
// 	db.AutoMigrate(&model.Books{})
// }
