package arguments

import (
	"fmt"
)

const pruneId ArgumentId = "prune"

func GetArgPruneDockerfile() Argument {
	return makeArgPrune(pruneDockerfile)
}

func GetArgPruneComposeFile() Argument {
	return makeArgPrune(pruneComposeFile)
}

func makeArgPrune(action func([]string)bool) Argument {
	return Argument{
		Id: pruneId,
		Action: action,
		ValidArgs: nil,
	}
}

func pruneDockerfile(args []string) bool {
	fmt.Println("Deletando todos Dockerfiles...")

	return true
}

func pruneComposeFile(args []string) bool {
	fmt.Println("Deletados todos compose files...")

	return true
}

