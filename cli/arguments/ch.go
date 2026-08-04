package arguments

import (
	"errors"
	"fmt"

	filem "didockerf/file"
	"didockerf/model"
	"didockerf/out"
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

	dockerfilesIdentifier := model.IdentifierStr(args[0])
	newIdentifier := model.IdentifierStr(args[1])

	errOnChange := filem.ChangeSavedDockerfileIdentifier(dockerfilesIdentifier, newIdentifier)
	if errOnChange != nil {
		out.FatalError(errOnChange)
		return true
	}

	return true
}

func checkIfChangeDockerfileArgsAreCorrect(args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("Change dockerfile requires 2 arguments, but received %d.", len(args)))
	}

	savedDockerfileIdentifierStr := model.IdentifierStr(args[0])
	_, err := filem.TransformSavedDockerfileIdentifierStrIntoDockerfile(savedDockerfileIdentifierStr)
	if err != nil {
		return errors.New("The saved dockerfile with the passed identifer doesnt't exist")
	}

	newIdentifierStr := model.IdentifierStr(args[1])
	if !model.IsIdentifierStrValid(newIdentifierStr) {
		return errors.New("New dockerfile's identifier is incorrectly formatted. It must be <name>:<tag>.")
	}
	return nil
}

func changeComposeFile(args []string) bool {
	fmt.Println("Alterando compose file...")

	return true
}
