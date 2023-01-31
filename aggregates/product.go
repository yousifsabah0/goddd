package aggregates

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/entity"
)

var (
	ErrMissingValues = errors.New("missing values")
)

type Product struct {
	product  *entity.Product
	price    float64
	quantity int
}

func NewProduct(title, description string, price float64) (Product, error) {
	if title == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		product: &entity.Product{
			Id:          uuid.New(),
			Title:       title,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p *Product) GetId() uuid.UUID {
	return p.product.Id
}

func (p *Product) GetProduct() *entity.Product {
	return p.product
}

func (p *Product) GetPrie() float64 {
	return p.price
}
