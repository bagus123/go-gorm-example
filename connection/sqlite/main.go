package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// driver sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "my_db.db")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connected")

	defer db.Close()
}
