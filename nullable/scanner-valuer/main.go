package main

import (
	"database/sql"
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
	Name         string     `json:"name"`
	Price        NullInt64  `json:"price" gorm:"DEFAULT:0"`
	FreeShipping bool       `json:"free_shipping"`
	FreeGift     bool       `json:"free_gift"`
	Description  NullString `json:"description" gorm:"DEFAULT:''"`
}

// extends
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
// implement MarshalJSON
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
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

	products := []Product{
		{Name: "Susu"},
		{Name: "Coklat"},
	}

	for _, product := range products {
		db.Create(&product)
	}

	db.Where(&Product{
		Name:        "Coklat",
		Price:       NullInt64{sql.NullInt64{Int64: 0, Valid: true}},
		Description: NullString{sql.NullString{String: "", Valid: true}}}).Find(&products)

	arrByte, _ = json.MarshalIndent(products, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

}
