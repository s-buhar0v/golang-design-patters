package menu

type MenuComponent interface {
	GetName() (string, error)
	GetDescription() (string, error)
	GetPrice() (float32, error)
	IsVegetarian() (bool, error)
	Print()
	AddComponent(mc MenuComponent) error
	GetChild(i int) (MenuComponent, error)
}
