package devices

import "fmt"

type VideoPlayer interface {
	On()
	InsertDisk(name string)
	ExtractDisk(name string)
	Play(name string)
	Stop(name string)
	Off()
}

type DVD struct{}

func NewDVD() VideoPlayer {
	return &DVD{}
}

func (d *DVD) On() {
	fmt.Println("Turn on dvd")
}
func (d *DVD) InsertDisk(name string) {
	fmt.Printf("Insert '%s' disk\n", name)
}

func (d *DVD) ExtractDisk(name string) {
	fmt.Printf("Extract '%s' disk\n", name)
}

func (d *DVD) Play(name string) {
	fmt.Printf("Play '%s'\n", name)
}

func (d *DVD) Stop(name string) {
	fmt.Printf("Stop '%s'\n", name)
}

func (d *DVD) Off() {
	fmt.Println("Turn off dvd")
}
