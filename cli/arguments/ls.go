package arguments

import (
	"fmt"

	filem "didockerf/file"
	"didockerf/out"
)

const listID ArgumentID = "ls"

func GetArgLsDockerfile() Argument {
	return makeArgLs(listDockerfiles)
}

func GetArgLsComposeFile() Argument {
	return makeArgLs(listComposeFiles)
}

func makeArgLs(action func([]string) bool) Argument {
	return Argument{
		ID:        listID,
		Action:    action,
		ValidArgs: nil,
	}
}

func listDockerfiles(args []string) bool {
	savedDockerfiles := filem.GetAllSavedDockerfiles()
	out.PrintDockerfilesInTable(savedDockerfiles)

	return true
}

func listComposeFiles(args []string) bool {
	fmt.Println("Listando compose files...")

	return true
}
