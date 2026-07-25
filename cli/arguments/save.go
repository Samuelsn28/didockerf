package arguments

import (
	"fmt"
	"errors"
	"strings"

	filem "didockerf/file"
	"didockerf/model"
	"didockerf/out"
)

const saveId ArgumentId = "save"

func GetArgSaveDockerfile() Argument {
	return makeArgSave(saveDockerfile)
}

func GetArgSaveComposeFile() Argument {
	return makeArgSave(saveComposeFile)
}

func makeArgSave(action func([]string)bool) Argument {
	return Argument{
		Id: saveId,
		Action: action,
		ValidArgs: nil,
	}
}

func saveDockerfile(args []string) bool {
	errOnArgs := checkIfSaveDockerfileArgsAreCorrect(args)
	if (errOnArgs != nil) {
		out.PrintError(errOnArgs)
		return true
	}

	dockerfilePath := args[0]
	saveIdentifierSplitted := splitSaveIdentifier(args[1])
	saveName := saveIdentifierSplitted[0]
	saveTag := "1.0"

	if (len(saveIdentifierSplitted) == 2) {
		saveTag = saveIdentifierSplitted[1]
	}

	dockerfile := model.Dockerfile{
		Name: saveName,
		Tag: saveTag,
		OriginPath: dockerfilePath,
	}

	fmt.Println("Salvando dockerfile...")

	filem.SaveDockerfile(dockerfile)

	return true
}

func checkIfSaveDockerfileArgsAreCorrect(args []string) error {
	if (len(args) != 2) {
		return errors.New(fmt.Sprintf("Save requires 2 arguments, but received %d.", len(args)))
	}

	dockerfilePath := args[0]
	if (!existDockerfileToSave(dockerfilePath)) { 
		return errors.New("Passed dockerfile doesn't exist.")
	}

	saveIdentifier := args[1]
	errInIdentifier := checkIfSaveIdentifierIsCorrectlyFormatted(saveIdentifier)
	if (errInIdentifier != nil) {
		return errInIdentifier
	}

	return nil
}

func existDockerfileToSave(dockerfilePath string) bool {
	return filem.FileExist(dockerfilePath)
}

func checkIfSaveIdentifierIsCorrectlyFormatted(saveIdentifier string) error {
	saveIdentifierSplitted := splitSaveIdentifier(saveIdentifier)

	// TODO: limitar caracteres que podem ser usados no nome (evitar problemas para indicar diretorio)
	
	if (len(saveIdentifierSplitted) <= 0 || len(saveIdentifierSplitted) > 2) {
		return errors.New("Dockerfile identifier must be <name> or <name>:<tag>")
	}
	return nil
}

func saveComposeFile(args []string) bool {
	fmt.Println("Salvando compose file...")

	return true
}

func splitSaveIdentifier(saveIdentifier string) []string {
	return strings.Split(saveIdentifier, ":")
}

