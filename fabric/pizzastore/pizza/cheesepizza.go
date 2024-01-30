package pizza

import (
	"fmt"

	"github.com/s-buhar0v/golang-design-patters/decorator/pizzastore/ingredients"
)

type ItalianCheesePizza struct{ *Pizza }

func NewItalianCheesePizza(i ingredients.IngredientsFabric) *ItalianCheesePizza {
	return &ItalianCheesePizza{
		Pizza: &Pizza{
			Name:              "Cheese",
			ingredientsFabric: i,
		},
	}
}

func (p *ItalianCheesePizza) Prepare() {
	fmt.Printf("Preparing %s ...\n", p.Pizza.Name)

	dough := p.Pizza.ingredientsFabric.CreateDough()
	fmt.Printf("Roll the %s ...\n", dough)

	sauce := p.Pizza.ingredientsFabric.CreateSauce()
	fmt.Printf("Add the %s sauce ...\n", sauce)

	cheese := p.Pizza.ingredientsFabric.CreateCheese()
	fmt.Printf("Add the %s cheese ...\n", cheese)
}

type AmericanCheesePizza struct{ *Pizza }

func NewAmericanCheesePizza(i ingredients.IngredientsFabric) *AmericanCheesePizza {
	return &AmericanCheesePizza{
		Pizza: &Pizza{
			Name:              "Cheese",
			ingredientsFabric: i,
		},
	}
}

func (p *AmericanCheesePizza) Prepare() {
	fmt.Printf("Preparing %s ...\n", p.Pizza.Name)

	dough := p.Pizza.ingredientsFabric.CreateDough()
	fmt.Printf("Roll the %s ...\n", dough)

	sauce := p.Pizza.ingredientsFabric.CreateSauce()
	fmt.Printf("Add the %s sauce ...\n", sauce)

	cheese := p.Pizza.ingredientsFabric.CreateCheese()
	fmt.Printf("Add the %s cheese ...\n", cheese)
}
