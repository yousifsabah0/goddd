package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/aggregates"
	"github.com/yousifsabah0/goddd/domain/customer"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregates.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregates.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregates.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return aggregates.Customer{}, customer.ErrNoCustomerFound
}

func (mr *MemoryRepository) Add(c aggregates.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregates.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exists %w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(c aggregates.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer does not exists %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()

	return nil
}
