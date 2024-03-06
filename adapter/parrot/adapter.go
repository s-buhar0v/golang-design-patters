package parrot

import (
	"fmt"

	"github.com/s-buhar0v/golang-design-patters/adapter/duck"
)

type MallardDuckAdapter struct {
	p Parrot
}

func NewMallardDuckAdapter(p Parrot) duck.Duck {
	return &MallardDuckAdapter{p: p}
}

func (m *MallardDuckAdapter) Fly() {
	m.p.Fly()
}

func (m *MallardDuckAdapter) Quack() {
	m.p.Sing()
	m.p.Talk()
}

func (m *MallardDuckAdapter) Name() {
	fmt.Println("I am rad macaw!")
}
