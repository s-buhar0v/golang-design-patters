package parrot

import "fmt"

// Parrot is base interface.
type Parrot interface {
	Fly()
	Sing()
	Talk()
}

type RedMacaw struct{}

func (m *RedMacaw) Fly() {
	fmt.Println("I can fly!")
}

func (m *RedMacaw) Sing() {
	fmt.Println("I can sing!")
}

func (m *RedMacaw) Talk() {
	fmt.Println("I can talk!")
}
