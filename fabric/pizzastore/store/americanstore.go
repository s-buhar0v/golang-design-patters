package store

import (
	"github.com/s-buhar0v/golang-design-patters/fabric/pizzastore/ingredients"
	"github.com/s-buhar0v/golang-design-patters/fabric/pizzastore/pizza"
)

type AmericanPizzaStore struct{ *pizzaStore }

func NewAmericanPizzaStore() IPizzaStore {
	i := &AmericanPizzaStore{}
	i.pizzaStore = &pizzaStore{
		createCheesePizza: i.createCheesePizza,
	}

	return i
}

func (i *AmericanPizzaStore) createCheesePizza() pizza.IPizza {
	return pizza.NewAmericanCheesePizza(&ingredients.AmericanIngredientsFabric{})
}
