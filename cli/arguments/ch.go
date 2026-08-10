package arguments

import (
	"errors"
	"fmt"

	"didockerf/out"
	savem "didockerf/savesManagement"
)

const changeId ArgumentID = "ch"

func GetArgChangeDockerfile() Argument {
	return makeArgChange(changeDockerfile)
}

func GetArgChangeComposeFile() Argument {
	return makeArgChange(changeComposeFile)
}

func makeArgChange(action func([]string) bool) Argument {
	return Argument{
		ID:        changeId,
		Action:    action,
		ValidArgs: nil,
	}
}

func changeDockerfile(args []string) bool {
	errOnArgs := checkIfChangeDockerfileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	dockerfileIdentifierStr := args[0]
	newIdentifierStr := args[1]

	errOnChange := savem.ChangeSavedDockerfileIdentifier(dockerfileIdentifierStr, newIdentifierStr)
	if errOnChange != nil {
		out.FatalError(errOnChange)
		return true
	}

	return true
}

func changeComposeFile(args []string) bool {
	errOnArgs := checkIfChangeComposeFileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	composeFileIdentifierStr := args[0]
	newIdentifierStr := args[1]

	errOnChange := savem.ChangeSavedComposeFileIdentifier(composeFileIdentifierStr, newIdentifierStr)
	if errOnChange != nil {
		out.FatalError(errOnArgs)
		return true
	}

	return true
}

func checkIfChangeDockerfileArgsAreCorrect(args []string) error {
	return checkIfChangeArgsAreCorrect(args)
}

func checkIfChangeComposeFileArgsAreCorrect(args []string) error {
	return checkIfChangeArgsAreCorrect(args)
}

func checkIfChangeArgsAreCorrect(args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("Change requires 2 arguments, but received %d.", len(args)))
	}
	return nil
}
