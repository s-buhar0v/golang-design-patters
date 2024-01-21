package beverage

import "fmt"

type Beverage interface {
	Cost() int
	GetDescription() string
}

type Espresso struct{}

func NewEspresso() Beverage {
	return &Espresso{}
}

func (e *Espresso) Cost() int {
	return 10
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

// Decorators

// Beverage with milk
type Milk struct {
	beverage Beverage
}

func WithMilk(b Beverage) Beverage {
	return &Milk{beverage: b}
}

func (m *Milk) Cost() int {
	return m.beverage.Cost() + 1
}

func (m *Milk) GetDescription() string {
	return fmt.Sprintf("%s, Milk", m.beverage.GetDescription())
}

// Vert big beverage
type KingSize struct {
	beverage Beverage
}

func WithKingSize(b Beverage) Beverage {
	return &KingSize{beverage: b}
}

func (k *KingSize) Cost() int {
	return k.beverage.Cost() * 2
}

func (k *KingSize) GetDescription() string {
	return fmt.Sprintf("%s, King size", k.beverage.GetDescription())
}
