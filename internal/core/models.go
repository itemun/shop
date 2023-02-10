package core

import "errors"

type Shop struct {
	ProductCount int
}

func newShop() *Shop {
	return &Shop{
		ProductCount: 100,
	}
}

func (s *Shop) sell() (int, error) {
	if s.ProductCount == 0 {
		return 0, errors.New("no product in the shop")
	}
	s.ProductCount--
	return s.ProductCount, nil
}

func (s *Shop) buy() int {
	s.ProductCount++
	return s.ProductCount
}
