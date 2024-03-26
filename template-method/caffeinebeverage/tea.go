package caffeinebeverage

import "fmt"

type Tea struct {
	*CaffeineBeverage
	addLemon bool
}

func NewTea(addLemon bool) ICaffeineBeverage {
	t := &Tea{addLemon: addLemon}
	c := &CaffeineBeverage{
		Brew:                    t.brew,                    // implementing abstract method
		AddCondiments:           t.addCondiments,           // implementing abstract method
		CustomerWantsCondiments: t.customerWantsCondiments, // implementing virtual method
	}
	t.CaffeineBeverage = c

	return t
}

func (t *Tea) brew() error {
	fmt.Println("Brewing strong tea ...")

	return nil
}

func (t *Tea) addCondiments() error {
	fmt.Println("Adding lemon ...")

	return nil
}

func (t *Tea) customerWantsCondiments() bool {
	return t.addLemon
}
