package main

import "github.com/s-buhar0v/golang-design-patters/template-method/caffeinebeverage"

var (
	tea          caffeinebeverage.ICaffeineBeverage
	teaWithLemon caffeinebeverage.ICaffeineBeverage
	coffee       caffeinebeverage.ICaffeineBeverage
)

func main() {
	tea = caffeinebeverage.NewTea(false)
	teaWithLemon = caffeinebeverage.NewTea(true)
	coffee = caffeinebeverage.NewCoffe()

	tea.Prepare()
	teaWithLemon.Prepare()
	coffee.Prepare()
}
