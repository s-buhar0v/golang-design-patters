package main

import (
	"fmt"

	"github.com/s-buhar0v/golang-design-patters/decorator/pizzastore/store"
)

func main() {
	italianPizzaStore := store.NewItalianPizzaStore()
	italianPizzaStore.OrderCheesePizza()

	fmt.Println()

	americanPizzaStore := store.NewAmericanPizzaStore()
	americanPizzaStore.OrderCheesePizza()
}
