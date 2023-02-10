package main

import (
	"errors"
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

func newShop() *Shop {
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

func (s *Shop) sell(id string) error {
	if product, ok := s.Products[id]; ok {
		if product.Quantity == 0 {
			return errors.New("no product in the shop")
		}
		product.Quantity-- //меняет количество товара у копии, а не у оригинала
		//s.Products[id].Quantity-- // не работает
		fmt.Println(product.Name, product.Quantity)
	} else {
		return fmt.Errorf("no product with id %s in the shop", id)
	}
	return nil
}

func (s *Shop) buy(id string) error {
	if product, ok := s.Products[id]; ok {
		product.Quantity++
	} else {
		return fmt.Errorf("no product with id %s in the shop", id)
	}
	return nil
}

func main() {
	shop := newShop()
	err := shop.sell("2")
	if err != nil {
		panic(err)
	}
	fmt.Println(shop)
}
