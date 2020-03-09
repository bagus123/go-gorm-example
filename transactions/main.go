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

func CreateProductManualTrx(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&Product{Name: "Minyak Goreng"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Product{Name: "Roti Keju"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// err = errors.New("Raise some error")
	// if err != nil {
	// 	return err
	// }

	return tx.Commit().Error

}

func CreateProduct(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&Product{Name: "Susu Kuda Liar", Price: 100000, FreeShipping: true}).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err = tx.Model(&Product{}).Where("name LIKE ?", "%Susu%").Update("name", "Coklat").Error; err != nil {
			return err
		}

		// err = errors.New("Raise some error")
		// if err != nil {
		// 	return err
		// }

		// return nil will commit
		return nil
	})
}

var (
	err     error
	db      *gorm.DB
	arrByte []byte
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

	err = CreateProduct(db)

	if err != nil {
		panic(err)
	}

	err = CreateProductManualTrx(db)
	if err != nil {
		panic(err)
	}

	var product Product
	db.Take(&product)

	arrByte, _ = json.MarshalIndent(product, "", "\t")
	fmt.Printf("Result 1 :\n %v \n", string(arrByte))

}
