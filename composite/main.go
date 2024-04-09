package main

import "github.com/s-buhar0v/golang-design-patters/composite/menu"

var (
	dailyMenu     menu.MenuComponent
	breakfastMenu menu.MenuComponent

	teaWithLemon menu.MenuComponent
	coffee       menu.MenuComponent
)

func main() {
	dailyMenu = menu.NewMenu("Daily", "It is our general menu")
	breakfastMenu = menu.NewMenu("Breakfast", "This menu is available till 13:00")
	dailyMenu.AddComponent(breakfastMenu)

	teaWithLemon = menu.NewMenuItem(
		"Tea",
		"Black tea with lemon",
		0.25,
		true,
	)
	dailyMenu.AddComponent(teaWithLemon)

	coffee = menu.NewMenuItem(
		"Coffee",
		"Delicious dark coffee",
		0.3,
		true,
	)
	breakfastMenu.AddComponent(coffee)

	dailyMenu.Print()
}
