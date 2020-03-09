package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Place belongs to Town
// it would be a place being part of a town with the foreign key on the place.
type Place struct {
	Id     int
	Name   string
	Town   Town
	TownId int //Foregin key
}

type Town struct {
	Id   int
	Name string
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

	db.DropTableIfExists(&Place{})
	db.DropTableIfExists(&Town{})

	db.CreateTable(&Place{})
	db.CreateTable(&Town{})

	t := Town{
		Name: "Jakarta",
	}

	p1 := Place{
		Name:   "Semanggi",
		TownId: 1,
	}

	p2 := Place{
		Name:   "Manggarai",
		TownId: 1,
	}

	p3 := Place{
		Name: "Cipondoh",
		Town: Town{
			Name: "Tangerang",
		},
	}

	p4 := Place{
		Name: "Cikupa",
		Town: Town{
			Name: "Tangerang",
		},
	}

	db.Save(&t)
	db.Save(&p1)
	db.Save(&p2)
	db.Save(&p3)
	db.Save(&p4)

	log.Println(p3)

	places := []Place{}
	db.Find(&places)
	for i := range places {
		db.Model(places[i]).Related(&places[i].Town)
	}

	fmt.Println(places)
	arrByte, _ = json.Marshal(places)
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

	var town Town
	db.Where(Town{Name: "Tangerang"}).FirstOrCreate(&town)
	db.Create(&Place{Name: "Cipondoh", TownId: town.Id})

	var place_ Place
	db.Where(&Place{Name: "Cipondoh"}).First(&place_)
	db.Model(place_).Related(&place_.Town)
	log.Println(place_)

	arrByte, _ = json.Marshal(place_)
	fmt.Printf("Find Result :\n %v \n", string(arrByte))

}
