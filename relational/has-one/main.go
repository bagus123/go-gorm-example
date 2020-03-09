package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Has one associations are associations where the foreign key for the one-to-one relation exists on the target model.
// User has one CreditCard, CreditCardID is the foreign key
type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
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

	// log mode if true, will info query SQL
	db.LogMode(true)

	db.DropTableIfExists(&CreditCard{})
	db.DropTableIfExists(&User{})

	db.CreateTable(&CreditCard{})
	db.CreateTable(&User{})

	//db.AutoMigrate(&CreditCard{}, &User{})

	var u User
	db.Where(&User{Name: "John Doe"}).FirstOrCreate(&u)

	u1 := &User{
		Name:       "Edward Robinson",
		CreditCard: CreditCard{Number: "0001"},
	}

	db.Create(u1)

	cc := CreditCard{Number: "0002", UserID: u.ID}
	db.Save(&cc)

	var user User
	db.First(&user)
	db.Model(user).Related(&user.CreditCard)

	arrByte, _ = json.MarshalIndent(user, "", "\t")
	fmt.Printf("Result 1 :\n %v \n", string(arrByte))

	var user2 User
	db.Preload("CreditCard").First(&user2, "name=?", "Edward Robinson")
	//db.Debug().Preload("CreditCard").First(&user2, "name=?", "Edward Robinson")
	arrByte, _ = json.MarshalIndent(user2, "", "\t")
	fmt.Printf("Result 2 :\n %v \n", string(arrByte))

	var user3 User
	db.Find(&user3, "name = ?", "John Doe")
	db.Model(&user3).Association("CreditCard").Find(&user3.CreditCard)
	arrByte, _ = json.MarshalIndent(user3, "", "\t")
	fmt.Printf("Result 3 :\n %v \n", string(arrByte))

	var users []User
	db.Find(&users)
	for _, user := range users {
		db.Model(&user).Association("CreditCard").Find(&user.CreditCard)
	}

	arrByte, _ = json.MarshalIndent(users, "", "\t")
	fmt.Printf("Result 4 :\n %v \n", string(arrByte))

}
