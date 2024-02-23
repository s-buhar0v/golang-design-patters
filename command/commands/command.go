package commands

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type DefaultCommand struct{}

func NewDefaultCommand() Command {
	return &DefaultCommand{}
}

func (d *DefaultCommand) Execute() {
	fmt.Println("Default command do nothing")
}

func (d *DefaultCommand) Undo() {
	fmt.Println("Default command do nothing")
}
