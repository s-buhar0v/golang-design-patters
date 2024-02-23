package commands

import "github.com/s-buhar0v/golang-design-patters/command/devices"

type TurnOnTVCommand struct {
	tv *devices.TV
}

func NewTurnOnTVCommand() Command {
	return &TurnOnTVCommand{}
}

func (t *TurnOnTVCommand) Execute() {
	t.tv.On()
}

func (t *TurnOnTVCommand) Undo() {
	t.tv.Off()
}

type TurnOffTVCommand struct {
	tv *devices.TV
}

func NewTurnOffTVCommand() Command {
	return &TurnOffTVCommand{}
}

func (t *TurnOffTVCommand) Execute() {
	t.tv.Off()
}

func (t *TurnOffTVCommand) Undo() {
	t.tv.On()
}
