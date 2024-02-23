package devices

import "fmt"

type TV struct{}

func (t *TV) On() {
	fmt.Println("Turn on TV")
}

func (t *TV) Off() {
	fmt.Println("Turn off TV")
}
