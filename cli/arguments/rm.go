package arguments

import (
	"fmt"
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
	fmt.Println("Removendo dockerfile...")

	return true
}

func removeComposeFile(args []string) bool {
	fmt.Println("Removendo compose file...")

	return true
}
