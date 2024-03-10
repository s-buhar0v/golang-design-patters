package main

import (
	"github.com/s-buhar0v/golang-design-patters/facade/devices"
	"github.com/s-buhar0v/golang-design-patters/facade/movieplayer"
)

func main() {
	movieName := "The Green Mile"
	moviePlayer := movieplayer.NewMoviePlayer()
	moviePlayer.SetDevices(devices.NewPlasmaTV(), devices.NewDVD())

	moviePlayer.Start(movieName)
	moviePlayer.Stop(movieName)
}
