package ingredients

type ItalianIngredientsFabric struct{}

func (i *ItalianIngredientsFabric) CreateDough() Dough {
	return "thin crust"
}

func (i *ItalianIngredientsFabric) CreateCheese() Cheese {
	return "mozarella"
}

func (i *ItalianIngredientsFabric) CreateSauce() Sauce {
	return "classic tomato"
}
