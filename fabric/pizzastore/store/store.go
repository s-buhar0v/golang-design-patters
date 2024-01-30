package store

import "github.com/s-buhar0v/golang-design-patters/decorator/pizzastore/pizza"

type IPizzaStore interface {
	createCheesePizza() pizza.IPizza
	OrderCheesePizza()
}

type pizzaStore struct {
	createCheesePizza func() pizza.IPizza
}

func (p *pizzaStore) OrderCheesePizza() {
	pizza := p.createCheesePizza()

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
}
