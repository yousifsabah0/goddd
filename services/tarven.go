package services

import (
	"log"

	"github.com/google/uuid"
)

type TarvenConfiguration func(ts *Tarven) error

type Tarven struct {
	OrderService *OrderService
}

func NewTarven(cfgs ...TarvenConfiguration) (*Tarven, error) {
	t := &Tarven{}

	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (t *Tarven) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("Bill the customer %0.0f", price)
	return nil
}

func WithOrderService(os *OrderService) TarvenConfiguration {
	return func(t *Tarven) error {
		t.OrderService = os
		return nil
	}
}
