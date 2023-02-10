package core

import (
	"fmt"
)

type Product struct {
	Name     string
	Price    int
	Quantity int
}

type Shop struct {
	Products map[string]Product
}

func NewShop() *Shop {
	return &Shop{
		Products: map[string]Product{
			"1": {
				Name:     "Phone",
				Price:    50000,
				Quantity: 100,
			},
			"2": {
				Name:     "TV",
				Price:    100000,
				Quantity: 50,
			},
			"3": {
				Name:     "Notebook",
				Price:    150000,
				Quantity: 30,
			},
		}}
}

func (s *Shop) Sell(id string) error {
	if product, ok := s.Products[id]; ok {
		if product.Quantity == 0 {
			return errProductNotFound
		}
		product.Quantity-- //меняет количество товара у копии, а не у оригинала
		//s.Products[id].Quantity-- // не работает
		fmt.Println(product.Name, product.Quantity)
	} else {
		return errProductNotFound
	}
	return nil
}

func (s *Shop) Buy(id string) error {
	if product, ok := s.Products[id]; ok {
		product.Quantity++
	} else {
		return errProductNotFound
	}
	return nil
}
