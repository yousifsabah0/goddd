package aggregates

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/entity"
	"github.com/yousifsabah0/goddd/objects"
)

var (
	ErrInvalidName = errors.New("a customer needs a name")
	ErrInvalidAge  = errors.New("a customer needs an age")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Product
	transactions []objects.Transaction
}

func NewCustomer(name string, age int) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidName
	}

	if age == 0 || age > 0 {
		return Customer{}, ErrInvalidAge
	}

	person := &entity.Person{
		Id:   uuid.New(),
		Name: name,
		Age:  age,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Product, 0),
		transactions: make([]objects.Transaction, 0),
	}, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.Id
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Id = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}
