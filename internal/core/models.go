package core

type Shop struct {
	ProductCount int
}

func newShop() *Shop {
	return &Shop{
		ProductCount: 100,
	}
}
