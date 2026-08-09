package arguments

import (
	"errors"
	"fmt"

	"didockerf/out"
	savem "didockerf/savesManagement"
)

const rmID ArgumentID = "rm"

func GetArgRmDockerfile() Argument {
	return makeArgRm(removeDockerfile)
}

func GetArgRmComposeFile() Argument {
	return makeArgRm(removeComposeFile)
}

func makeArgRm(action func([]string) bool) Argument {
	return Argument{
		ID:        rmID,
		Action:    action,
		ValidArgs: nil,
	}
}

func removeDockerfile(args []string) bool {
	errOnArgs := checkIfArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	savedDockerfileIdentifierStr := args[0]

	errOnRemove := savem.RemoveSavedDockerfile(savedDockerfileIdentifierStr)
	if errOnRemove != nil {
		out.FatalError(errOnRemove)
		return true
	}

	return true
}

func checkIfArgsAreCorrect(args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("Remove dockerfile requires 1 argument, but received %d.", len(args)))
	}

	return nil
}

func removeComposeFile(args []string) bool {
	fmt.Println("Removendo compose file...")

	return true
}
