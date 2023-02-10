package core

type Product struct {
	Name     string
	Price    int
	Quantity int
}

type Shop struct {
	Products map[string]*Product
}

func NewShop() *Shop {
	return &Shop{
		Products: map[string]*Product{
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
	product, ok := s.Products[id]
	if !ok || product.Quantity == 0 {
		return errProductNotFound
	}
	product.Quantity--
	return nil
}

func (s *Shop) Buy(id string) error {
	product, ok := s.Products[id]
	if !ok || product.Quantity == 0 {
		return errProductNotFound
	}
	product.Quantity++
	return nil
}
