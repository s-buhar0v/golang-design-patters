package remotecontrol

import (
	"errors"

	"github.com/s-buhar0v/golang-design-patters/command/commands"
)

type RemoteControl struct {
	maxSlots int

	onCommands  map[int]commands.Command
	offCommands map[int]commands.Command

	lastCommand commands.Command
}

func NewRemoteControl(maxSlots int) *RemoteControl {
	onCommands := map[int]commands.Command{}
	offCommands := map[int]commands.Command{}

	for i := 0; i < maxSlots; i++ {
		onCommands[i] = commands.NewDefaultCommand()
		offCommands[i] = commands.NewDefaultCommand()
	}

	return &RemoteControl{
		maxSlots: maxSlots,

		onCommands:  onCommands,
		offCommands: offCommands,
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand commands.Command, offCommand commands.Command) error {
	if slot > r.maxSlots {
		return errors.New("limit reached")
	}

	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand

	return nil
}

func (r *RemoteControl) PressOnButton(slot int) {
	r.onCommands[slot].Execute()
	r.lastCommand = r.onCommands[slot]
}

func (r *RemoteControl) PressOffButton(slot int) {
	r.offCommands[slot].Execute()
	r.lastCommand = r.offCommands[slot]
}

func (r *RemoteControl) PressUndoButton() {
	r.lastCommand.Undo()
}
