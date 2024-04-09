package menu

import (
	"errors"
	"fmt"
)

type Menu struct {
	name        string
	description string
	items       []MenuComponent
}

func NewMenu(
	name string,
	description string,
) MenuComponent {
	return &Menu{
		name:        name,
		description: description,
	}
}

func (m *Menu) GetName() (string, error) {
	return m.name, nil
}

func (m *Menu) GetDescription() (string, error) {
	return m.description, nil
}

func (m *Menu) GetPrice() (float32, error) {
	return 0, errors.New("unsupported method")
}

func (m *Menu) IsVegetarian() (bool, error) {
	return false, errors.New("unsupported method")
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s\n", m.name, m.description)

	for _, i := range m.items {
		i.Print()
	}
}

func (m *Menu) AddComponent(mc MenuComponent) error {
	m.items = append(m.items, mc)

	return nil
}

func (m *Menu) GetChild(i int) (MenuComponent, error) {
	return m.items[i], nil
}
