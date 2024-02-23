package commands

import "github.com/s-buhar0v/golang-design-patters/command/devices"

type TurnOnHeatingCommand struct {
	heating *devices.Heating
}

func NewTurnOnHeatingCommand() Command {
	return &TurnOnHeatingCommand{}
}

func (t *TurnOnHeatingCommand) Execute() {
	t.heating.Enable()
}

func (t *TurnOnHeatingCommand) Undo() {
	t.heating.Disalbe()
}

type TurnOffHeatingCommand struct {
	h *devices.Heating
}

func NewTurnOffHeatingCommand() Command {
	return &TurnOffHeatingCommand{}
}

func (t *TurnOffHeatingCommand) Execute() {
	t.h.Disalbe()
}

func (t *TurnOffHeatingCommand) Undo() {
	t.h.Enable()
}
