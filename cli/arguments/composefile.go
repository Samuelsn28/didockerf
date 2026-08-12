package arguments

import "didockerf/out/printers"

const composeFileCategoryID ArgumentID = "composefile"

func GetSubCommandComposeFile() Argument {
	return Argument{
		ID:        composeFileCategoryID,
		Action:    changeCategoryToComposeFile,
		ValidArgs: getValidArgsForComposeFileCategory(),
	}
}

func changeCategoryToComposeFile(args []string) bool {
	if len(args) == 0 {
		showComposeFileHelpMsg()
		return true
	}

	return false
}

func getValidArgsForComposeFileCategory() map[ArgumentID]Argument {
	ls := GetArgLsComposeFile()
	save := GetArgSaveComposeFile()
	get := GetArgGetComposeFile()
	rm := GetArgRmComposeFile()
	ch := GetArgChangeComposeFile()
	prune := GetArgPruneComposeFile()

	return map[ArgumentID]Argument{
		ls.ID:    ls,
		save.ID:  save,
		get.ID:   get,
		rm.ID:    rm,
		ch.ID:    ch,
		prune.ID: prune,
	}
}

func showComposeFileHelpMsg() {
	helpMsg := `
Usage:	didockerf composefile <command>

Commands:
  save		Save a new compose file
  ls		List all saved compose files
  get		Retrieve a saved compose file
  ch		Change identifier of a saved compose file
  rm		Remove a saved compose file
  prune		Remove all saved compose files

	`

	printers.PrintHelpMessage(helpMsg)
}
