package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// driver postgre
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=postgres sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connected")

	defer db.Close()
}
