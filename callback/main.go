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

func WithFreeShipping(db *gorm.DB) *gorm.DB {
	return db.Where("free_shipping = ?", true)
}

func WithFreeGift(db *gorm.DB) *gorm.DB {
	return db.Where("free_gift = ?", true)
}

func updateCreated(scope *gorm.Scope) {
	if scope.HasColumn("Description") {
		err := scope.SetColumn("Description", "set by system")
		if err != nil {
			panic(err)
		}
	}
}

func updateCreatedAfter(scope *gorm.Scope) {
	if scope.HasColumn("Description") {
		err := scope.SetColumn("Description", "set by system B")
		if err != nil {
			panic(err)
		}
	}
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

	db.DropTableIfExists(&Product{})
	db.AutoMigrate(&Product{})

	db.Callback().Create().Before("gorm:create").Register("satu", updateCreated)
	db.Callback().Create().After("gorm:create").Register("dua", updateCreatedAfter)

	products := []Product{
		{Name: "Susu", Price: 1000, FreeShipping: true, FreeGift: true},
		{Name: "Coklat", Price: 1000, FreeShipping: false, FreeGift: false},
	}

	for _, product := range products {
		db.Create(&product)
	}

	// Add function of scope
	db.Scopes(WithFreeShipping, WithFreeGift).Find(&products)

	arrByte, _ = json.MarshalIndent(products, "", "\t")
	fmt.Printf("Result 1 :\n %v \n", string(arrByte))

}
