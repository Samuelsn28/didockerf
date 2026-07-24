package arguments

import (
	"fmt"
	"strings"

	filem "didockerf/file"
	"didockerf/model"
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
	if (!areSaveDockerfileArgsCorrect(args)) {
		return true
	}

	dockerfilePath := args[0]
	saveIdentifierSplitted := splitSaveIdentifier(args[1])
	saveName := saveIdentifierSplitted[0]
	saveVersion := "1.0"

	if (len(saveIdentifierSplitted) == 2) {
		saveVersion = saveIdentifierSplitted[1]
	}

	dockerfile := model.Dockerfile{
		Name: saveName,
		Version: saveVersion,
		OriginPath: dockerfilePath,
	}

	fmt.Println("Salvando dockerfile...")

	filem.SaveDockerfile(dockerfile)

	return true
}

func areSaveDockerfileArgsCorrect(args []string) bool {
	if (len(args) != 2) {
		fmt.Println("[!] Esperava 2 argumentos, mas não veio essa quantidade!")
		return false
	}

	dockerfilePath := args[0]
	if (!existDockerfileToSave(dockerfilePath)) { 
		return false
	}

	// TODO: limitar caracteres que podem ser usados no nome (evitar problemas para indicar diretorio)
	saveIdentifier := args[1]
	if (!isSaveIdentifierCorrectlyFormatted(saveIdentifier)) {
		return false
	}

	return true
}

func existDockerfileToSave(dockerfilePath string) bool {
	return filem.FileExist(dockerfilePath)
}

func isSaveIdentifierCorrectlyFormatted(saveIdentifier string) bool {
	saveIdentifierSplitted := splitSaveIdentifier(saveIdentifier)

	if (len(saveIdentifierSplitted) < 0 || len(saveIdentifierSplitted) > 2) {
		return false
	}
	return true
}

func saveComposeFile(args []string) bool {
	fmt.Println("Salvando compose file...")

	return true
}

func splitSaveIdentifier(saveIdentifier string) []string {
	return strings.Split(saveIdentifier, ":")
}

