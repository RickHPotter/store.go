package models

type Product struct {
	id          uint16
	name        string
	description string
	price       float64
	inStock     int
}

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
