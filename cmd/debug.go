package cmd

import (
	"github.com/fatih/color"
)

type Debug struct {
	message string
}

func NewDebug(message string) *Debug {
	return &Debug{
		message: message,
	}
}

func (d *Debug) Success() {
	// A newline will be appended automatically
	color.Green("Prints %s in green.", d.message)
}

func (d *Debug) Error() {
	// A newline will be appended automatically
	color.Red("Prints %s in red.", d.message)
}
