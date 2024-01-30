package store

import (
	"github.com/s-buhar0v/golang-design-patters/decorator/pizzastore/ingredients"
	"github.com/s-buhar0v/golang-design-patters/decorator/pizzastore/pizza"
)

type ItalianPizzaStore struct{ *pizzaStore }

func NewItalianPizzaStore() IPizzaStore {
	i := &ItalianPizzaStore{}
	i.pizzaStore = &pizzaStore{
		createCheesePizza: i.createCheesePizza,
	}

	return i
}

func (i *ItalianPizzaStore) createCheesePizza() pizza.IPizza {
	return pizza.NewItalianCheesePizza(&ingredients.ItalianIngredientsFabric{})
}
