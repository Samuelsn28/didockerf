package arguments

import (
	"fmt"
)

const listId ArgumentId = "ls"

func GetArgLsDockerfile() Argument {
	return makeArgLs(listDockerfiles)
}

func GetArgLsComposeFile() Argument {
	return makeArgLs(listComposeFiles)
}

func makeArgLs(action func([]string)bool) Argument {
	return Argument{
		Id: listId,
		Action: action,
		ValidArgs: nil,
	}
}

func listDockerfiles(args []string) bool {
	fmt.Println("Listando dockerfiles...")

	return true
}

func listComposeFiles(args []string) bool {
	fmt.Println("Listando compose files...")

	return true
}

