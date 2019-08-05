package entities

import "fmt"

type Product struct {
	Id       int64   `json:"id"` //json:"id" id adalah nama custom yang ditampilkan di result
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"id_category"`
}

type ProductEdit struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"id_category"`
}

func (product Product) toString() string {
	return fmt.Sprintf("id: %d\nname: %d\nprice: %0.1\nquantity: %d\ncategory: %d", product.Id, product.Name, product.Price, product.Price, product.Quantity, product.Category)
}

func (productEdit ProductEdit) toString() string {
	return fmt.Sprintf("name: %d\nprice: %0.1\nquantity: %d\ncategory: %d", productEdit.Name, productEdit.Price, productEdit.Price, productEdit.Quantity, productEdit.Category)
}
