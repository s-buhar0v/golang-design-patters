package duck

import "fmt"

// Duck is base interface.
type Duck interface {
	Fly()
	Quack()
	Name()
}

// Duck is base interface.
type MallardDuck struct{}

func (m *MallardDuck) Fly() {
	fmt.Println("I can fly!")
}

func (m *MallardDuck) Quack() {
	fmt.Println("I can quack!")
}

func (m *MallardDuck) Name() {
	fmt.Println("I am mallard duck!")
}
