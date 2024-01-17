package duck

import "fmt"

type FlyBehaviuor interface {
	Fly()
}

type WingFly struct{}

func (w *WingFly) Fly() {
	fmt.Println("I'm flying with wings!")
}

type RocketFly struct{}

func (w *RocketFly) Fly() {
	fmt.Println("I'm flying via rocket!")
}

type NoFly struct{}

func (w *NoFly) Fly() {
	fmt.Println("I can't fly :(")
}
