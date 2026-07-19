package arguments

import (
	"fmt"
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
	fmt.Println("Alterando dockerfile...")

	return true
}

func changeComposeFile(args []string) bool {
	fmt.Println("Alterando compose file...")

	return true
}

