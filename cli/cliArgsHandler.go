package cli

import (
	"os"

	argPkg "didockerf/cli/arguments"
	"didockerf/out"
)

var remainingArgs []string = os.Args[1:]

func Operate(validArgs map[argPkg.ArgumentID]argPkg.Argument) {
	if len(remainingArgs) == 0 {
		remainingArgs = append(remainingArgs, "--help")
	}

	currentPassedArg := remainingArgs[0]
	arg, isArgValid := validArgs[argPkg.ArgumentID(currentPassedArg)]
	if !isArgValid {
		out.PrintFatalError("Invalid argument.")
		return
	}

	remainingArgs = remainingArgs[1:]

	finishExecution := arg.Action(remainingArgs)

	if finishExecution {
		return
	}

	validArgs = arg.ValidArgs
	Operate(validArgs)
}
