package storage

import "github.com/itemun/shop/internal/core"

type Storage struct {
	Products map[string]*core.Product
}

// NewStorage returns a new Storage
func NewStorage() *Storage {
	return &Storage{
		Products: make(map[string]*core.Product),
	}
}

// Create adds a new product
func (s *Storage) Create(id string, product *core.Product) error {
	if _, ok := s.Products[id]; ok {
		return core.ErrProductExists
	}
	s.Products[id] = product
	return nil
}

// Update updates a product
func (s *Storage) Update(id string, product *core.Product) error {
	if _, ok := s.Products[id]; !ok {
		return core.ErrProductNotFound
	}
	s.Products[id] = product // check how to update product: every field must be updated?

	return nil
}

// Delete deletes a product
func (s *Storage) Delete(id string) error {
	if _, ok := s.Products[id]; ok {
		delete(s.Products, id)
		return nil
	}
	return core.ErrProductNotFound

}

// GetAll returns a list of products
func (s *Storage) GetAll() []*core.Product {
	products := make([]*core.Product, 0, len(s.Products))
	for _, v := range s.Products {
		products = append(products, v)
	}
	return products
}

// GetByID returns a product by id
func (s *Storage) GetByID(id string) (*core.Product, error) {
	if product, ok := s.Products[id]; ok {
		return product, nil
	}
	return nil, core.ErrProductNotFound
}
