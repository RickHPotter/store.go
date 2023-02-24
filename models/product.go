package models

import (
	"github.com/RickHPotter/web_app_alura_course/db"
)

type Product struct {
	id          uint16
	name        string
	description string
	price       float64
	inStock     int
}

/*
DEV INTERFACE
*/

/*
USER INTERFACE
*/

func NewProduct(id uint16, name string, description string, price float64, inStock int) *Product {
	product := Product{id, name, description, price, inStock}
	return &product
}

func (product Product) GetName() string {
	return product.name
}

func (product Product) GetDescription() string {
	return product.description
}

func (product Product) GetPrice() float64 {
	return product.price
}

func (product Product) GetInStock() int {
	return product.inStock
}

func ListAll() []Product {
	db := db.Connect()
	defer db.Close()

	query, e := db.Query("select * from products")
	if e != nil {
		panic(e.Error())
	}

	products := []Product{}

	for query.Next() {
		var id uint16
		var stock int
		var name, description string
		var price float64

		e = query.Scan(&id, &name, &description, &price, &stock)
		if e != nil {
			panic(e.Error())
		}

		p := NewProduct(id, name, description, price, stock)
		products = append(products, *p)
	}
	return products
}
