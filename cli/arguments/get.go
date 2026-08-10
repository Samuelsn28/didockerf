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
	errOnArgs := checkIfGetDockerfileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	savedDockerfileIdentifierStr := args[0]
	copyToPath := getDestinationPath(args)

	errOnCopyDockerfile := savem.CopySavedDockerfileTo(savedDockerfileIdentifierStr, copyToPath)
	if errOnCopyDockerfile != nil {
		out.FatalError(errOnCopyDockerfile)
		return true
	}

	return true
}

func getSavedComposeFile(args []string) bool {
	errOnArgs := checkIfGetComposeFileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	savedComposeFileIdentifierStr := args[0]
	copyToPath := getDestinationPath(args)

	errOnCopyComposeFile := savem.CopySaveComposeFileTo(savedComposeFileIdentifierStr, copyToPath)
	if errOnCopyComposeFile != nil {
		out.FatalError(errOnCopyComposeFile)
		return true
	}

	return true
}

func checkIfGetDockerfileArgsAreCorrect(args []string) error {
	return checkIfGetArgsAreCorrect(args)
}

func checkIfGetComposeFileArgsAreCorrect(args []string) error {
	return checkIfGetArgsAreCorrect(args)
}

func checkIfGetArgsAreCorrect(args []string) error {
	if len(args) < 1 || len(args) > 2 {
		return fmt.Errorf("Get requires 1 or 2 arguments, but received %d", len(args))
	}

	return nil
}

func getDestinationPath(args []string) string {
	if len(args) == 1 {
		return "."
	}
	return args[1]
}
