package arguments

import (
	"fmt"
	"strings"

	filem "didockerf/file"
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
	// Verifica argumentos
	if (len(args) != 2) {
		fmt.Println("[!] Esperava 2 argumentos, mas não veio essa quantidade!")
		return true
	}

	dockerfilePath := args[0]

	if !filem.FileExist(dockerfilePath) { 
		return true
	}

	// TODO: limitar caracteres que podem ser usados no nome (evitar problemas para indicar diretorio)
	saveIdentifier := strings.Split(args[1], ":")

	if (len(saveIdentifier) > 2 || len(saveIdentifier) < 0) {
		fmt.Println("[!] Argumento Nome:Versão mal formatado")
		return true
	}

	saveName := saveIdentifier[0]
	saveVersion := "1.0"

	if (len(saveIdentifier) == 2) {
		saveVersion = saveIdentifier[1]
	}

	fmt.Println("Salvando dockerfile...")

	filem.CreateDockerfileSave(dockerfilePath, saveName, saveVersion)

	return true
}

func saveComposeFile(args []string) bool {
	fmt.Println("Salvando compose file...")

	return true
}

