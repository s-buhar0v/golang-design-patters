package devices

import "fmt"

type Heating struct{}

func (t *Heating) Enable() {
	fmt.Println("Enable heating")
}

func (t *Heating) Disalbe() {
	fmt.Println("Disalbe heating")
}
