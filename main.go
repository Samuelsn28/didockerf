package main

import (
	"didockerf/cli"
	pkgArgs "didockerf/cli/arguments"
)

func main() {
	determineSubCommand()
}

func determineSubCommand() {
	subCommands := pkgArgs.GetSubCommands()

	cli.Operate(subCommands)


}



