package devices

import "fmt"

type TV interface {
	On()
	Off()
}

type PlasmaTV struct{}

func NewPlasmaTV() *PlasmaTV {
	return &PlasmaTV{}
}

func (p *PlasmaTV) On() {
	fmt.Println("Turn on plasma")
}

func (p *PlasmaTV) Off() {
	fmt.Println("Turn off plasma")
}
