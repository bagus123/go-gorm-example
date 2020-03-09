package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" sql:"index"`

	// set as pointer because can set as nill / null
	// nullable value
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

type Product struct {
	BaseModel
	Name         string `json:"name"`
	Price        int    `json:"price"`
	FreeShipping bool   `json:"free_shipping"`
	FreeGift     bool   `json:"free_gift"`
	Description  string `json:"description"`
}

type People struct {
	BaseModel
	Name string `json:"name"`
}

var (
	arrByte []byte
	db      *gorm.DB
	err     error
)

func main() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=postgres sslmode=disable")

	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	defer db.Close()

	db.LogMode(true)
	db.DropTableIfExists(&Product{})
	db.DropTableIfExists(&People{})
	db.AutoMigrate(&Product{}, &People{})

	db.Unscoped().Delete(&People{})

	products := []Product{
		{Name: "Susu", Price: 12000},
		{Name: "Coklat", Price: 7000},
	}

	for _, product := range products {
		db.Create(&product)
	}

	peoples := []People{
		{Name: "John Doe"},
		{Name: "Edward"},
	}

	for _, people := range peoples {
		db.Create(&people)
	}

	var product Product
	db.Raw("SELECT name FROM products WHERE name = ?", "Coklat").Scan(&product)

	arrByte, _ = json.MarshalIndent(product, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

	// RUNNING STORED PROCEDURE OR FUNCTION
	result := struct {
		Z int
	}{}
	db.Raw("SELECT * FROM math_add(?,?)", 1, 2).Scan(&result)
	fmt.Println(result)

	result2 := struct {
		W int
		Z int
	}{}
	db.Raw("SELECT * FROM math_add2(?,?)", 1, 2).Scan(&result2)
	fmt.Println(result2)

	result3 := []People{}
	db.Raw("SELECT * from getpeople()").Scan(&result3)
	fmt.Println(result3)
	arrByte, _ = json.MarshalIndent(result3, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

}
