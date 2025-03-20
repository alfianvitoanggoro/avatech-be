package main

import "github.com/alfianvitoanggoro/avatech/cmd"

func main() {
	debug := cmd.NewDebug("Hello World")
	debug.Success()
	debug.Error()
}
