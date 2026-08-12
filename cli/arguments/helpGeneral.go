package arguments

import (
	"didockerf/out/printers"
)

const helpID ArgumentID = "help"

func GetSubCommandHelpGeneral() Argument {
	return makeArgHelp(printGeneralHelp)
}

func makeArgHelp(action func([]string) bool) Argument {
	return Argument{
		ID:        helpID,
		Action:    action,
		ValidArgs: nil,
	}
}

func printGeneralHelp(args []string) bool {
	helpMsg := `
Usage:	didockerf <command> 

Dockerfiles commands:
  dfile		Operate in dockerfile mode

Compose files commands:
  composefile	Operate in compose file mod

Others commands:
  help		Shows this message

`

	printers.PrintHelpMessage(helpMsg)

	return true
}
