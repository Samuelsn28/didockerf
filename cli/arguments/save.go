package arguments

import (
	"fmt"
)

const saveId string = "save"

func GetArgSaveDockerfile() Argument {
	return makeArgSave(saveDockerfile)
}

func GetArgSaveComposeFile() Argument {
	return makeArgSave(saveComposeFile)
}

func makeArgSave(action func([]string)bool) Argument {
	return Argument{
		Id: ArgumentId(saveId),
		Action: action,
		ValidArgs: nil,
	}
}

func saveDockerfile(args []string) bool {
	fmt.Println("Salvando Dockerfile...")

	return true
}

func saveComposeFile(args []string) bool {
	fmt.Println("Salvando compose file...")

	return true
}

