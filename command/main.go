package main

import (
	"github.com/s-buhar0v/golang-design-patters/command/commands"
	"github.com/s-buhar0v/golang-design-patters/command/remotecontrol"
)

const (
	maxSlots = 1
)

func main() {
	r := remotecontrol.NewRemoteControl(maxSlots)
	r.SetCommand(
		1,
		commands.NewTurnOnAllDevicesCommand(
			[]commands.Command{
				commands.NewTurnOnTVCommand(),
				commands.NewTurnOnHeatingCommand(),
			},
		),
		commands.NewTurnOffAllDevicesCommand(
			[]commands.Command{
				commands.NewTurnOffTVCommand(),
				commands.NewTurnOffHeatingCommand(),
			},
		))

	r.PressOnButton(1)
	// Prints
	// Turn on TV
	// Enable heating

	r.PressOffButton(1)
	// Prints
	// Turn off TV
	// Disalbe heating

	r.PressUndoButton()
	// Prints
	// Turn on TV
	// Enable heating
}
