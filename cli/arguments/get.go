package arguments

import (
	"fmt"

	"didockerf/out"
	savem "didockerf/savesManagement"
)

const getID ArgumentID = "get"

func GetArgGetDockerfile() Argument {
	return makeArgGet(getSavedDockerfile)
}

func GetArgGetComposeFile() Argument {
	return makeArgGet(getSavedComposeFile)
}

func makeArgGet(action func([]string) bool) Argument {
	return Argument{
		ID:        getID,
		Action:    action,
		ValidArgs: nil,
	}
}

func getSavedDockerfile(args []string) bool {
	errOnArgs := checkIfPassedArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	savedDockerfileIdentifierStr := args[0]

	var copyToPath string
	if len(args) == 2 && savem.DirExist(args[1]) {
		copyToPath = args[1]
	} else {
		copyToPath = "."
	}

	errOnCopyDockerfile := savem.CopySavedDockerfileTo(savedDockerfileIdentifierStr, copyToPath)
	if errOnCopyDockerfile != nil {
		out.FatalError(errOnCopyDockerfile)
		return true
	}

	return true
}

func checkIfPassedArgsAreCorrect(args []string) error {
	if len(args) < 1 || len(args) > 2 {
		return fmt.Errorf("Get a dockerfile requires 1 or 2 arguments, but received %d", len(args))
	}

	return nil
}

func getSavedComposeFile(args []string) bool {
	return true
}
