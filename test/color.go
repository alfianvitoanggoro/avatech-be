package test

import "github.com/alfianvitoanggoro/avatech/cmd"

func Color() {
	debug := cmd.NewDebug("Hello World")
	debug.Success()
	debug.Error()
}
