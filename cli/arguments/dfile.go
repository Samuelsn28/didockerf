package arguments

import "didockerf/out/printers"

const dockerfileCategoryID ArgumentID = "dfile"

func GetSubCommandDfile() Argument {
	return Argument{
		ID:        dockerfileCategoryID,
		Action:    changeCategoryToDockerfile,
		ValidArgs: getValidArgsForDockerfileCategory(),
	}
}

func changeCategoryToDockerfile(args []string) bool {
	if len(args) == 0 {
		showDfileHelpMsg()
		return true
	}

	return false
}

func getValidArgsForDockerfileCategory() map[ArgumentID]Argument {
	ls := GetArgLsDockerfile()
	save := GetArgSaveDockerfile()
	get := GetArgGetDockerfile()
	rm := GetArgRmDockerfile()
	ch := GetArgChangeDockerfile()
	prune := GetArgPruneDockerfile()

	return map[ArgumentID]Argument{
		ls.ID:    ls,
		save.ID:  save,
		get.ID:   get,
		rm.ID:    rm,
		ch.ID:    ch,
		prune.ID: prune,
	}
}

func showDfileHelpMsg() {
	helpMsg := `
Usage:	didockerf dfile <command>

Commands:
  save		Save a new dockerfile
  ls		List all saved dockerfiles
  get		Retrieve a saved dockerfile
  ch		Change identifier of a saved dockerfile
  rm		Remove a saved dockerfile
  prune		Remove all saved dockerfiles

	`

	printers.PrintHelpMessage(helpMsg)
}
