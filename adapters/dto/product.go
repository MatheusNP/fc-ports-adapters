package dto

import "github.com/MatheusNP/fc-ports-adapters/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	if _, err := product.IsValid(); err != nil {
		return &application.Product{}, err
	}

	return product, nil
}
