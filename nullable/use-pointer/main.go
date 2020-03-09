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
	Name         string  `json:"name"`
	Price        *int    `json:"price" gorm:"DEFAULT:0"`
	FreeShipping bool    `json:"free_shipping"`
	FreeGift     bool    `json:"free_gift"`
	Description  *string `json:"description" gorm:"DEFAULT:''"`
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
	db.CreateTable(&Product{})

	products := []Product{
		{Name: "Susu"},
		{Name: "Coklat"},
	}

	for _, product := range products {
		db.Create(&product)
	}

	price := 0
	desc := ""
	db.Where(&Product{Name: "Coklat",
		Price:       &price,
		Description: &desc}).Find(&products)
	// SELECT * FROM "products"  WHERE "products"."deleted_at" IS NULL AND (("products"."name" = 'Coklat') AND ("products"."price" = 0) AND ("products"."description" = ''))

	arrByte, _ = json.MarshalIndent(products, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

}
