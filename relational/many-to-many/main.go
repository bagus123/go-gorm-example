package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tidwall/pretty"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" sql:"index"`

	// set as pointer because can set as nill / null
	// nullable value
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	BaseModel
	Name      string     `json:"name"`
	Languages []Language `json:"language" gorm:"many2many:user_languages;"`
}

type Language struct {
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

	db.DropTableIfExists(&Language{})
	db.DropTableIfExists(&User{})

	db.CreateTable(&Language{})
	db.CreateTable(&User{})

	languages := []Language{{Name: "Indonesia"}, {Name: "Inggris"}}
	db.Create(&User{Name: "John Doe", Languages: languages})

	var user User
	db.First(&user)
	db.Model(user).Related(&user.Languages, "Languages")
	log.Println(user)

	arrByte, _ = json.Marshal(user)

	// for prettier can use tab, and comment lib pretty
	//arrByte, _ = json.MarshalIndent(user, "", "\t")

	// prettier using lib
	arrByte = pretty.Pretty(arrByte)

	fmt.Printf("Find Result :\n %v \n", string(arrByte))
}
