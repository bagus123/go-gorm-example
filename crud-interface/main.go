package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Table name is the pluralized version of struct name.
// to be 'products' table
type Product struct {
	gorm.Model // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `products`
	Code       string
	Price      uint
}

var (
	arrByte []byte
	db      *gorm.DB
	err     error
)

func main() {

	fmt.Println("run")

	db, err = gorm.Open("sqlite3", "my_db.db")
	//db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=postgres sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	// defer fmt.Println("close")
	// will execute after in the end of line code (after execute db.Delete)
	defer db.Close()

	// Migrate the schema
	db.LogMode(true)
	db.DropTableIfExists(&Product{})
	db.AutoMigrate(&Product{})

	// delete by query
	db.Delete(Product{}, "created_at LIKE ?", "%-%")

	// Create Row 1
	db.Create(&Product{Code: "0001", Price: 1000})

	// Create Row 2
	db.Create(&Product{Code: "0002", Price: 1200})

	// Create Row 3
	db.Create(&Product{Code: "0003", Price: 3000})

	p1 := &Product{Code: "0003", Price: 3000}
	err = db.Create(p1).Error

	fmt.Println(p1.ID)

	if err != nil {
		panic(err)
	}

	result := db.Create(&Product{Code: "0003", Price: 3000})

	if result.Error != nil {
		panic(err)
	}

	var products []Product

	// Get all records
	db.Find(&products)
	//db.Where("SQL", "valie").Find()
	arrByte, _ = json.MarshalIndent(products, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

	fmt.Println("")
	fmt.Println("")

	// Read
	var product Product
	db.First(&product, 1)                  // find product with id 1
	db.First(&product, "code = ?", "0001") // find product with code 0001
	arrByte, _ = json.MarshalIndent(product, "", "\t")
	fmt.Printf("First Result :\n %v \n", string(arrByte))

	fmt.Println("")
	fmt.Println("")

	// Get first matched record
	db.Where("price = ?", 12000).Where("code = ?", "jinzhu").First(&product)

	//// SELECT * FROM products WHERE name = 'jinzhu' limit 1;
	arrByte, _ = json.MarshalIndent(product, "", "\t")
	fmt.Printf("First Result : \n %v \n", string(arrByte))

	fmt.Println("")
	fmt.Println("")

	// Get all matched records
	db.Where("code = ?", "0001").Find(&products)
	//// SELECT * FROM products WHERE name = '0001';
	arrByte, _ = json.MarshalIndent(products, "", "\t")
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

	fmt.Println("")
	fmt.Println("")

	// Where Struct
	db.Where(&Product{Code: "0001", Price: 1000}).First(&product)
	// SELECT * FROM products WHERE code = "0001" AND price = 1000 LIMIT 1;
	arrByte, _ = json.MarshalIndent(product, "", "\t")
	fmt.Printf("First Result : \n %v \n", string(arrByte)) // output single object product

	fmt.Println("")
	fmt.Println("")

	// Where Struct
	db.Where(&Product{Code: "0001", Price: 1000}).Find(&products)
	// SELECT * FROM products WHERE code = "0001" AND price = 1000;
	arrByte, _ = json.MarshalIndent(products, "", "\t")
	fmt.Printf("Find Result : \n %v \n", string(arrByte)) // output Array of product

	fmt.Println("")
	fmt.Println("")

	// Update - update product's price to 99999
	// Price (field model) or price (field in table) can do
	db.Model(&product).Update("Price", 99999)

	// // Update also can do by this code
	// product.Code = "0002"
	// db.Save(&product)

	// // Delete - delete product
	// db.Delete(&product)
}
