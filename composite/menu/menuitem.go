package menu

import (
	"errors"
	"fmt"
)

type MenuItem struct {
	name         string
	description  string
	price        float32
	isVegetarian bool
}

func NewMenuItem(
	name string,
	description string,
	price float32,
	isVegetarian bool,
) MenuComponent {
	return &MenuItem{
		name:         name,
		description:  description,
		price:        price,
		isVegetarian: isVegetarian,
	}
}

func (m *MenuItem) GetName() (string, error) {
	return m.name, nil
}

func (m *MenuItem) GetDescription() (string, error) {
	return m.description, nil
}

func (m *MenuItem) GetPrice() (float32, error) {
	return m.price, nil
}

func (m *MenuItem) IsVegetarian() (bool, error) {
	return m.isVegetarian, nil
}

func (m *MenuItem) Print() {
	fmt.Printf("%s, %s, %.2f\n", m.name, m.description, m.price)
}

func (m *MenuItem) AddComponent(mc MenuComponent) error {
	return errors.New("unsupported method")
}

func (m *MenuItem) GetChild(i int) (MenuComponent, error) {
	return nil, errors.New("unsupported method")
}
