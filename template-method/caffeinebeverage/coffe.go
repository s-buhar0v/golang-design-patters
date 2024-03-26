package caffeinebeverage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coffee struct {
	*CaffeineBeverage
}

func NewCoffe() ICaffeineBeverage {
	coffee := &Coffee{}
	c := &CaffeineBeverage{
		Brew:                    coffee.brew,                    // implementing abstract method
		AddCondiments:           coffee.addCondiments,           // implementing abstract method
		CustomerWantsCondiments: coffee.customerWantsCondiments, // implementing virtual method
	}
	coffee.CaffeineBeverage = c

	return coffee
}

func (t *Coffee) brew() error {
	fmt.Println("Brewing coffee ...")

	return nil
}

func (t *Coffee) addCondiments() error {
	fmt.Println("Adding sugar ...")

	return nil
}

func (t *Coffee) customerWantsCondiments() bool {
	fmt.Println("Do you want to add sugar?")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil {
		return false
	}
	stringInput := strings.ToLower(string(input))

	if strings.ToLower(stringInput) == "yes" || strings.ToLower(stringInput) == "y" {
		return true
	}

	return false
}
