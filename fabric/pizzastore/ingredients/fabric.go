package ingredients

type IngredientsFabric interface {
	CreateDough() Dough
	CreateCheese() Cheese
	CreateSauce() Sauce
}
