package arguments

import (
	"errors"
	"fmt"

	"didockerf/out"
	"didockerf/out/printers"
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
	if len(args) == 0 {
		showChangeDockerfileHelpMsg()
		return true
	}

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
	if len(args) == 0 {
		showChangeComposeFileHelpMsg()
		return true
	}

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

func showChangeDockerfileHelpMsg() {
	helpMsg := `
Usage:	didockerf dfile ch <old name>:<old tag> <new name>:<new tag>

	`

	printers.PrintHelpMessage(helpMsg)
}

func showChangeComposeFileHelpMsg() {
	helpMsg := `
Usage:	didockerf composefile ch <old name>:<old tag> <new name>:<new tag>

	`

	printers.PrintHelpMessage(helpMsg)
}
