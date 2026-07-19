package arguments

import (
	"fmt"
)

const rmId ArgumentId = "rm"

func GetArgRmDockerfile() Argument {
	return makeArgRm(removeDockerfile)
}

func makeArgRm(action func([]string)bool) Argument {
	return Argument{
		Id: rmId,
		Action: action,
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


