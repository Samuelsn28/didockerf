package arguments

import (
	"errors"
	"fmt"

	"didockerf/out"
	"didockerf/out/printers"
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
	if len(args) == 0 {
		showRemoveDockerfileHelpMsg()
		return true
	}

	errOnArgs := checkIfRemoveDockerfileArgsAreCorrect(args)
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

func removeComposeFile(args []string) bool {
	if len(args) == 0 {
		showRemoveComposeFileHelpMsg()
		return true
	}

	errOnArgs := checkIfRemoveComposeFileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	savedComposeFileIdentifierStr := args[0]

	errOnRemove := savem.RemoveSavedComposeFile(savedComposeFileIdentifierStr)
	if errOnRemove != nil {
		out.FatalError(errOnRemove)
		return true
	}

	return true
}

func checkIfRemoveDockerfileArgsAreCorrect(args []string) error {
	return checkIfRemoveArgsAreCorrect(args)
}

func checkIfRemoveComposeFileArgsAreCorrect(args []string) error {
	return checkIfRemoveArgsAreCorrect(args)
}

func checkIfRemoveArgsAreCorrect(args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("Remove requires 1 argument, but received %d.", len(args)))
	}

	return nil
}

func showRemoveDockerfileHelpMsg() {
	helpMsg := `
Usage:	didockerf dfile rm <name>:<tag>

	`

	printers.PrintHelpMessage(helpMsg)
}

func showRemoveComposeFileHelpMsg() {
	helpMsg := `
Usage:	didockerf composefile rm <name>:<tag>

	`

	printers.PrintHelpMessage(helpMsg)
}
