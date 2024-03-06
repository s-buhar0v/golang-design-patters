package main

import (
	"fmt"

	"github.com/s-buhar0v/golang-design-patters/adapter/duck"
	"github.com/s-buhar0v/golang-design-patters/adapter/parrot"
)

func whatCanDuck(d duck.Duck) {
	d.Name()
	d.Quack()
	d.Fly()
}

func main() {
	d := &duck.MallardDuck{}
	p := parrot.NewMallardDuckAdapter(&parrot.RedMacaw{})

	whatCanDuck(d)
	fmt.Println()
	whatCanDuck(p)
}
