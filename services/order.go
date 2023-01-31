package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/aggregates"
	"github.com/yousifsabah0/goddd/domain/customer"
	"github.com/yousifsabah0/goddd/domain/customer/memory"
	"github.com/yousifsabah0/goddd/domain/product"

	productMemory "github.com/yousifsabah0/goddd/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func (o *OrderService) CreateOrder(customerId uuid.UUID, productsIds []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerId)
	if err != nil {
		return 0, err
	}

	var products []aggregates.Product
	var price float64

	for _, id := range productsIds {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, nil
		}
		products = append(products, p)
		price += p.GetPrie()
	}

	log.Printf("Customer %v orderd %v", c.GetId(), len(products))

	return price, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregates.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := productMemory.New()
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}
