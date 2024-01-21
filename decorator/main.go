package main

import (
	"fmt"

	"github.com/s-buhar0v/golang-design-patters/decorator/beverage"
)

func main() {
	e := beverage.NewEspresso()
	e = beverage.WithMilk(e)
	e = beverage.WithKingSize(e)

	c := e.Cost()
	d := e.GetDescription()

	fmt.Printf("%s, %d$\n", d, c) // prin Espresso, Milk, King size, 22$
}
