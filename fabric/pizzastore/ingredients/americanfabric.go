package ingredients

type AmericanIngredientsFabric struct{}

func (i *AmericanIngredientsFabric) CreateDough() Dough {
	return "crust"
}

func (a *AmericanIngredientsFabric) CreateCheese() Cheese {
	return "cheddar"
}

func (i *AmericanIngredientsFabric) CreateSauce() Sauce {
	return "spicy tomato"
}
