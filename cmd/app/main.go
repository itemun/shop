package main

import (
	"fmt"
	"github.com/itemun/shop/internal/core"
)

func main() {
	shop := core.NewShop()
	err := shop.Sell("2")
	if err != nil {
		panic(err)
	}
	fmt.Println(shop)
}
