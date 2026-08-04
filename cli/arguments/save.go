package arguments

import (
	"errors"
	"fmt"
	"strings"

	filem "didockerf/file"
	"didockerf/model"
	"didockerf/out"
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
	errOnArgs := checkIfSaveDockerfileArgsAreCorrect(args)
	if errOnArgs != nil {
		out.FatalError(errOnArgs)
		return true
	}

	dockerfilePath := args[0]
	saveIdentifier, errCreateIdentifier := model.CreateIdentifierFromStr(model.IdentifierStr(args[1]))

	if errCreateIdentifier != nil {
		out.PrintFatalError("The passed identifier is invalid.")
		return true
	}

	dockerfile := model.CreateDockerfile(
		saveIdentifier,
		dockerfilePath,
	)

	filem.SaveDockerfile(dockerfile)

	return true
}

func checkIfSaveDockerfileArgsAreCorrect(args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("Save requires 2 arguments, but received %d.", len(args)))
	}

	dockerfilePath := args[0]
	if !existDockerfileToSave(dockerfilePath) {
		return errors.New("Passed dockerfile doesn't exist.")
	}

	saveIdentifier := model.IdentifierStr(args[1])
	if !model.IsIdentifierStrValid(saveIdentifier) {
		return errors.New("Passed identifier is invalid.")
	}

	return nil
}

func existDockerfileToSave(dockerfilePath string) bool {
	return filem.FileExist(dockerfilePath)
}

func saveComposeFile(args []string) bool {
	fmt.Println("Salvando compose file...")

	return true
}

func splitSaveIdentifier(saveIdentifier string) []string {
	return strings.Split(saveIdentifier, ":")
}
