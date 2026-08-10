package arguments

import (
	"didockerf/out/printers"
	savem "didockerf/savesManagement"
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
	savedDockerfilesInfos := savem.GetAllSavedDockerfilesInfos()
	printers.PrintSavedDockerfilesInfos(savedDockerfilesInfos)

	return true
}

func listComposeFiles(args []string) bool {
	savedComposeFilesInfos := savem.GetAllSavedComposeFilesInfos()
	printers.PrintSavedComposeFilesInfos(savedComposeFilesInfos)

	return true
}
