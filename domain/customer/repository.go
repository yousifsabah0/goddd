package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/aggregates"
)

var (
	ErrNoCustomerFound     = errors.New("customer not found")
	ErrFailedToAddCustomer = errors.New("failed to add a new customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(id uuid.UUID) (aggregates.Customer, error)
	Add(c aggregates.Customer) error
	Update(c aggregates.Customer) error
}
