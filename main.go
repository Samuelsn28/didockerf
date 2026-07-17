package main

import (
	"didockerf/cli"
	pkgArgs "didockerf/cli/arguments"
)

var alfa int = 25

func main() {
	determineSubCommand()

	alfa = 10
}

func determineSubCommand() {
	subCommands := pkgArgs.GetSubCommands()

	cli.Operate(subCommands)


}



