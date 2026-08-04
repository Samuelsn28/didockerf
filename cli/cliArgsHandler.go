package cli

import (
	"fmt"
	"os"

	argPkg "didockerf/cli/arguments"
)

var remainingArgs []string = os.Args[1:]

func Operate(validArgs map[argPkg.ArgumentID]argPkg.Argument) {
	if len(remainingArgs) == 0 {
		fmt.Println("Argumentos vazio. Encerrando...")
		return
	}

	fmt.Println(remainingArgs)

	currentPassedArg := remainingArgs[0]
	arg, isArgValid := validArgs[argPkg.ArgumentID(currentPassedArg)]
	if !isArgValid {
		// Error: comando inválido
		fmt.Println("Erro: argumento inválido")
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
