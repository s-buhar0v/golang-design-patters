package duck

import "fmt"

type QuackBehaviuor interface {
	Quack()
}

type DefaultQuack struct{}

func (w *DefaultQuack) Quack() {
	fmt.Println("Quack!")
}

type LoudQuack struct{}

func (w *LoudQuack) Quack() {
	fmt.Println("QUACK!")
}

type NoQuack struct{}

func (w *NoQuack) Quack() {
	fmt.Println("I can't quack :(")
}
