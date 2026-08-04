package arguments

import (
	"fmt"
)

const pruneID ArgumentID = "prune"

func GetArgPruneDockerfile() Argument {
	return makeArgPrune(pruneDockerfile)
}

func GetArgPruneComposeFile() Argument {
	return makeArgPrune(pruneComposeFile)
}

func makeArgPrune(action func([]string) bool) Argument {
	return Argument{
		ID:        pruneID,
		Action:    action,
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
