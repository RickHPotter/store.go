package models

import (
	"fmt"

	"github.com/RickHPotter/web_app_alura_course/db"
	"github.com/lib/pq"
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

func (product Product) _insert() {
	INSERT := `
	INSERT INTO PRODUCTS 
	(NAME, DESCRIPTION, PRICE, STOCK) 
	VALUES 
	($1, $2, $3, $4);`

	db := db.Connect()
	insert, e := db.Prepare(INSERT)
	if e != nil {
		panic(e.Error())
	}

	insert.Exec(product.name, product.description, product.price, product.inStock)

	defer db.Close()
}

func (product Product) _update() {
	UPDATE := `
	UPDATE PRODUCTS 
	SET NAME = $1, 
	DESCRIPTION = $2, 
	PRICE = $3, 
	STOCK = $4 
	WHERE CODPROD = $5;`

	db := db.Connect()
	update, e := db.Prepare(UPDATE)
	if e != nil {
		panic(e.Error())
	}
	fmt.Println(product.id)

	update.Exec(product.name, product.description, product.price, product.inStock, product.id)

	defer db.Close()
}

func _delete(id string) {
	DELETE := `
	DELETE FROM PRODUCTS 
	WHERE CODPROD = $1;`

	db := db.Connect()
	delete, e := db.Prepare(DELETE)
	if e != nil {
		panic(e.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

/*
USER INTERFACE
*/

func NewProduct(id uint16, name string, description string, price float64, inStock int) *Product {
	product := Product{id, name, description, price, inStock}
	return &product
}

func (product Product) GetId() uint16 {
	return product.id
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

	query, e := db.Query(
		`SELECT *
		FROM PRODUCTS
		ORDER BY CODPROD`)
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

func ListSome(where ...string) []Product {
	db := db.Connect()
	defer db.Close()

	query, e := db.Query(
		`SELECT *
		FROM PRODUCTS
		WHERE CODPROD = ANY($1)
		ORDER BY CODPROD`, pq.Array(where))
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

func (product Product) Insert() {
	// no rules to append here yet
	product._insert()
}

func (product Product) Update() {
	// no rules to append here yet
	product._update()
}

func Delete(id string) {
	// no rules to append here yet
	_delete(id)
}
