package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/aggregates"
	"github.com/yousifsabah0/goddd/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregates.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregates.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregates.Product, error) {
	var products []aggregates.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregates.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}

	return aggregates.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(p aggregates.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetId()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[p.GetId()] = p
	return nil
}

func (mpr *MemoryProductRepository) Update(p aggregates.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetId()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[p.GetId()] = p
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)
	return nil
}
