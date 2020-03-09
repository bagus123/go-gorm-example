package main

import (
	"encoding/json"
	"fmt"
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

type Customer struct {
	BaseModel
	CustomerName string
	Contacts     []Contact `json:"contacts"`
}

type Contact struct {
	BaseModel
	PhoneNo    string `json:"phone_no"`
	CustomerID int    `json:"customer_id"`
}

var (
	arrByte []byte
	db      *gorm.DB
	err     error
)

func main() {

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=postgres sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.LogMode(true)
	db.DropTableIfExists(&Contact{}, &Customer{})
	db.AutoMigrate(&Customer{}, &Contact{})

	c1 := Customer{CustomerName: "John", Contacts: []Contact{
		{PhoneNo: "0001"},
		{PhoneNo: "0002"}}}

	c2 := Customer{CustomerName: "Martin", Contacts: []Contact{
		{PhoneNo: "0003"},
		{PhoneNo: "0004"}}}

	c3 := Customer{CustomerName: "Ryan", Contacts: []Contact{
		{PhoneNo: "0005"},
		{PhoneNo: "0006"}}}

	c4 := Customer{CustomerName: "Edward", Contacts: []Contact{
		{PhoneNo: "0007"},
		{PhoneNo: "0008"}}}

	db.Create(&c1)
	db.Create(&c2)
	db.Create(&c3)
	db.Create(&c4)

	var customers []Customer
	var customer Customer
	//db.Debug().Where("customer_name=?", "Martin").Preload("Contacts").Find(&customer)
	db.Preload("Contacts").Find(&customers)

	arrByte, _ = json.MarshalIndent(customers, "", "\t")
	fmt.Printf("Result 1 :\n %v \n", string(arrByte))

	//Update
	db.Model(&Contact{}).Where("id=?", 3).Update("phone_no", "0100")

	db.Where("customer_name=?", "Martin").Preload("Contacts").Find(&customer)
	arrByte, _ = json.MarshalIndent(customer, "", "\t")
	fmt.Printf("Result 3 :\n %v \n", string(arrByte))

	//Count Associations
	contactCount := db.Model(&customer).Association("Contacts").Count()
	fmt.Printf("Result 4 :\n count costumer : %v \n", contactCount)

	//Clear Associations
	db.Model(&customer).Association("Contacts").Clear()

	//Delete
	db.Where("customer_name=?", customer.CustomerName).Delete(&customer)

	//Find All
	db.Preload("Contacts").Find(&customers)
	arrByte, _ := json.MarshalIndent(customers, "", "\t")
	fmt.Printf("Result 5 :\n %v \n", string(arrByte))
}
