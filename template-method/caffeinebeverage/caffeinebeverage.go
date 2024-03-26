package caffeinebeverage

import (
	"errors"
	"fmt"
)

type ICaffeineBeverage interface {
	Prepare()
	boilWater()
	pourInCup()

	brew() error
	addCondiments() error
	customerWantsCondiments() bool
}

type CaffeineBeverage struct {
	Brew                    func() error // abstract method
	AddCondiments           func() error // abstract method
	CustomerWantsCondiments func() bool  // virtual method
}

func NewCaffeineBeverage() *CaffeineBeverage {
	return &CaffeineBeverage{
		Brew:                    func() error { return errors.New("method must be implemented at child class") },
		AddCondiments:           func() error { return errors.New("method must be implemented at child class") },
		CustomerWantsCondiments: func() bool { return false },
	}
}

func (c *CaffeineBeverage) Prepare() {
	c.boilWater()
	c.Brew()
	if c.CustomerWantsCondiments() {
		c.AddCondiments()
	}
	c.pourInCup()
}

func (c *CaffeineBeverage) boilWater() {
	fmt.Println("Boiling water ...")
}

func (c *CaffeineBeverage) pourInCup() {
	fmt.Println("Pouring bevearage in the cup ...")
}
