package main

import (
	"fmt"
	"github.com/itemun/shop/internal/core"
)

func main() {
	shop := core.NewShop()
	err := shop.Buy("2")

	if err != nil {
		fmt.Println(err.Error()) // Зачем .Error()?
	}
	for id, product := range shop.Products {
		fmt.Printf("%s: %v\n", id, product)
	}
}

//Questions for Kostya:
//13 строка main
//Как пользуешься гитом в Голэнд?

//TODO
// Implement methods Sell & Buy
// In-memory cache в storage
// Handlers: get & post
// Tests
// Database
// Decorate DB
// Authorization/Authentication
