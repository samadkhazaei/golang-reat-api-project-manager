package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	database, err := gorm.Open("sqlite3", "./database/gorm.db")
	if err != nil {
		fmt.Printf(err.Error())
		panic("failed to connect database")
	}

	database.AutoMigrate(&Project{}, &Client{}, &User{})

	DB = database
}
