package commands

type TurnOnAllDevicesCommand struct {
	onCommands []Command
}

func NewTurnOnAllDevicesCommand(onCommands []Command) Command {
	return &TurnOnAllDevicesCommand{
		onCommands: onCommands,
	}
}

func (t *TurnOnAllDevicesCommand) Execute() {
	for _, c := range t.onCommands {
		c.Execute()
	}
}

func (t *TurnOnAllDevicesCommand) Undo() {
	for _, c := range t.onCommands {
		c.Undo()
	}
}

type TurnOffAllDevicesCommand struct {
	offCommands []Command
}

func NewTurnOffAllDevicesCommand(offCommands []Command) Command {
	return &TurnOffAllDevicesCommand{
		offCommands: offCommands,
	}
}

func (t *TurnOffAllDevicesCommand) Execute() {
	for _, c := range t.offCommands {
		c.Execute()
	}
}

func (t *TurnOffAllDevicesCommand) Undo() {
	for _, c := range t.offCommands {
		c.Undo()
	}
}
