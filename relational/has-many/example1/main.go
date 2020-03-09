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

// User has many CreditCards, UserID is the foreign key
type User struct {
	BaseModel
	Name        string       `json:"name"`
	CreditCards []CreditCard `json:"credit_cards"`
}

type CreditCard struct {
	BaseModel
	Number string `json:"name"`
	UserID uint   `json:"user_id"`
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

	db.DropTableIfExists(&CreditCard{})
	db.DropTableIfExists(&User{})

	db.CreateTable(&CreditCard{})
	db.CreateTable(&User{})

	var u User
	db.Where(&User{Name: "John Doe"}).FirstOrCreate(&u)

	cc1 := CreditCard{Number: "001", UserID: u.ID}
	cc2 := CreditCard{Number: "002", UserID: u.ID}

	db.Save(&cc1)
	db.Save(&cc2)

	var user User
	db.First(&user)
	db.Model(user).Related(&user.CreditCards)
	log.Println(user)

	arrByte, _ = json.MarshalIndent(user, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))
}
