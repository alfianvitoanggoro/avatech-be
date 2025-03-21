package debug

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug CLI",
	Long:  `An Debugging tools.`,
}

func def(message any) {
	var data any
	switch x := (message).(type) {
	default:
		data = x
	case any:
		b, _ := json.MarshalIndent(x, "", "    ")
		data = string(b)
	}
	fmt.Printf("%+v\n", data)
}

func success(scope string, message any) {
	status := color.New(color.BgHiGreen).Sprint(" SUCCESS ")
	scp := color.New(color.FgGreen).Sprint(scope)
	fmt.Printf("%s %s\n", status, scp)
	def(message)
}

func error(scope string, message any) {
	status := color.New(color.BgHiRed).Sprint(" ERROR ")
	scp := color.New(color.FgRed).Sprint(scope)
	fmt.Printf("%s %s\n", status, scp)
	def(message)
}
