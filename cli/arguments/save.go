package arguments

import (
	"errors"
	"fmt"

	"didockerf/out"
	"didockerf/out/printers"
	savem "didockerf/savesManagement"
)

const saveID ArgumentID = "save"

func GetArgSaveDockerfile() Argument {
	return makeArgSave(saveDockerfile)
}

func GetArgSaveComposeFile() Argument {
	return makeArgSave(saveComposeFile)
}

func makeArgSave(action func([]string) bool) Argument {
	return Argument{
		ID:        saveID,
		Action:    action,
		ValidArgs: nil,
	}
}

func saveDockerfile(args []string) bool {
	if len(args) == 0 {
		showSaveDockerfileHelpMsg()
		return true
	}

	errOnArgs := checkIfSaveDockerfileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	dockerfilePath := args[0]
	saveIdentifierStr := args[1]

	errOnSave := savem.SaveDockerfile(saveIdentifierStr, dockerfilePath)
	if errOnArgs != nil {
		out.FatalError(errOnSave)
		return true
	}

	return true
}

func saveComposeFile(args []string) bool {
	if len(args) == 0 {
		showSaveComposeFileHelpMsg()
		return true
	}

	errOnArgs := checkIfSaveComposeFileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	composeFilePath := args[0]
	saveIdentifierStr := args[1]

	errOnSave := savem.SaveComposeFile(saveIdentifierStr, composeFilePath)
	if errOnSave != nil {
		out.FatalError(errOnSave)
		return true
	}

	return true
}

func checkIfSaveDockerfileArgsAreCorrect(args []string) error {
	errOnArgsGeneral := checkIfSaveArgsAreCorrect(args)
	if errOnArgsGeneral != nil {
		return errOnArgsGeneral
	}

	dockerfilePath := args[0]
	if !existFileToSave(dockerfilePath) {
		return errors.New("Passed dockerfile doesn't exist.")
	}

	return nil
}

func checkIfSaveComposeFileArgsAreCorrect(args []string) error {
	errOnArgsGeneral := checkIfSaveArgsAreCorrect(args)
	if errOnArgsGeneral != nil {
		return errOnArgsGeneral
	}

	composeFilePath := args[0]
	if !existFileToSave(composeFilePath) {
		return errors.New("Passed compose file doesn't exist.")
	}

	return nil
}

func checkIfSaveArgsAreCorrect(args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("Save requires 2 arguments, but received %d.", len(args)))
	}

	return nil
}

func existFileToSave(filePath string) bool {
	return savem.FileExist(filePath)
}

func showSaveDockerfileHelpMsg() {
	helpMsg := `
Usage:	didockerf dfile save <dockerfile path> <name>:<name>

	`

	printers.PrintHelpMessage(helpMsg)
}

func showSaveComposeFileHelpMsg() {
	helpMsg := `
Usage:	didockerf composefile save <compose file path> <name>:<name>

	`

	printers.PrintHelpMessage(helpMsg)
}
