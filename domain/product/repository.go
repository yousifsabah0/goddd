package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/aggregates"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregates.Product, error)
	GetByID(id uuid.UUID) (aggregates.Product, error)
	Add(p aggregates.Product) error
	Update(p aggregates.Product) error
	Delete(id uuid.UUID) error
}
