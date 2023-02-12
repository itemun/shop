package core

import (
	"context"
	"fmt"
)

type ProductStorage interface {
	IncreaseProductCount(ctx context.Context, productID string, delta int) error
	DecreaseProductCount(ctx context.Context, productID string, delta int) error
}

type Product struct {
	Name     string
	Price    int
	Quantity int
}

type Shop struct {
	productStorage ProductStorage
}

func NewShop(storage ProductStorage) *Shop {
	return &Shop{
		productStorage: storage,
	}
}

func (s *Shop) Sell(ctx context.Context, id string) error {
	err := s.productStorage.DecreaseProductCount(ctx, id, 1)
	if err != nil {
		return fmt.Errorf("can't decrease product count: %w", err)
	}

	return nil
}

func (s *Shop) Buy(ctx context.Context, id string) error {
	err := s.productStorage.IncreaseProductCount(ctx, id, 1)
	if err != nil {
		return fmt.Errorf("can't increase product count: %w", err)
	}

	return nil
}
