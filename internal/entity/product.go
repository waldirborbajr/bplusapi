package entity

import "github.com/google/uuid"

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id, name string, price float64) *Product {
	return &Product{
		ID:    uuid.new().String(),
		Name:  name,
		Price: price,
	}
}
