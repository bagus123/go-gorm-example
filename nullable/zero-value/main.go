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
	db.AutoMigrate(&Product{})

	// zero values use default value
	// all fields having a zero value, like 0, '', false or other zero values,
	// won’t be saved into the database but will use its default value
	products := []Product{
		{Name: "Susu"},
		{Name: "Coklat"},
	}

	for _, product := range products {
		db.Create(&product)
	}

	db.Where(&Product{Name: "Coklat", Price: 0, Description: ""}).Find(&products)
	// query only "products"."name" = 'Coklat'
	// SELECT * FROM "products"  WHERE "products"."deleted_at" IS NULL AND (("products"."name" = 'Coklat'))
	// GORM will only query with those fields has non-zero value, that means
	// if your field’s value is 0, '', false or other zero values,
	// it won’t be used to build query conditions

	arrByte, _ = json.MarshalIndent(products, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

}
