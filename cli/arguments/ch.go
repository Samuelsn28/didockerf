package arguments

import (
	"errors"
	"fmt"

	filem "didockerf/file"
	"didockerf/out"
)

const changeId ArgumentId = "ch"

func GetArgChangeDockerfile() Argument {
	return makeArgChange(changeDockerfile)
}

func GetArgChangeComposeFile() Argument {
	return makeArgChange(changeComposeFile)
}

func makeArgChange(action func([]string)bool) Argument {
	return Argument{
		Id: changeId,
		Action: action,
		ValidArgs: nil,
	}
}

func changeDockerfile(args []string) bool {
	errOnArgs := checkIfChangeDockerfileArgsAreCorrect(args)
	if (errOnArgs != nil) {
		out.PrintError(errOnArgs)
		return true
	}

	dockerfilesIdentifier := args[0]
	newIdentifier := args[1]

	errOnChange := filem.ChangeSavedDockerfileIdentifier(dockerfilesIdentifier, newIdentifier)
	if (errOnChange != nil){
		out.PrintError(errOnChange)
		return true
	}

	return true
}

func checkIfChangeDockerfileArgsAreCorrect(args []string) error {
	if (len(args) != 2) {
		return errors.New(fmt.Sprintf("Change dockerfile requires 2 arguments, but received %d.", len(args)))
	}

	savedDockerfileIdentifier := args[0]
	_, err := filem.TransformSavedDockerfileIdentifierIntoDockerfile(savedDockerfileIdentifier)
	if (err != nil) {
		return errors.New("The saved dockerfile with the passed identifer doesnt't exist")
	}

	newIdentifier := args[1]
	if (!filem.IsIdentifierCorrectlyFormatted(newIdentifier)) {
		return errors.New("New dockerfile's identifier is incorrectly formatted. It must be <name>:<tag>.")
	}
	return nil
}

func changeComposeFile(args []string) bool {
	fmt.Println("Alterando compose file...")

	return true
}


