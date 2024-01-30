package pizza

import (
	"fmt"

	"github.com/s-buhar0v/golang-design-patters/decorator/pizzastore/ingredients"
)

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	Name string

	Dough    ingredients.Dough
	Cheese   ingredients.Cheese
	Toppings []ingredients.Topping
	Sauce    ingredients.Sauce

	Prepare func()

	ingredientsFabric ingredients.IngredientsFabric
}

func (p *Pizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350 ...")
}

func (p *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices ...")
}

func (p *Pizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box ...")
}
