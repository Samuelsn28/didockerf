package arguments

import (
	"errors"
	"fmt"

	"didockerf/out"
	savem "didockerf/savesManagement"
)

const pruneID ArgumentID = "prune"

func GetArgPruneDockerfile() Argument {
	return makeArgPrune(pruneDockerfile)
}

func GetArgPruneComposeFile() Argument {
	return makeArgPrune(pruneComposeFile)
}

func makeArgPrune(action func([]string) bool) Argument {
	return Argument{
		ID:        pruneID,
		Action:    action,
		ValidArgs: nil,
	}
}

func pruneDockerfile(args []string) bool {
	errOnArgs := checkIfPruneArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
	}

	errOnPrune := savem.RemoveAllSavedDockerfiles()
	if errOnPrune != nil {
		out.FatalError(errOnPrune)
	}

	return true
}

func checkIfPruneArgsAreCorrect(args []string) error {
	if len(args) != 0 {
		return errors.New("Prune command don't need any arguments.")
	}

	return nil
}

func pruneComposeFile(args []string) bool {
	fmt.Println("Deletados todos compose files...")

	return true
}
