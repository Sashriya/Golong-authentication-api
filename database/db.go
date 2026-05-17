package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	DB = db
}