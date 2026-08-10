package arguments

const composeFileCategoryID ArgumentID = "composefile"

func GetSubCommandComposeFile() Argument {
	return Argument{
		ID:        composeFileCategoryID,
		Action:    changeCategoryToComposeFile,
		ValidArgs: getValidArgsForComposeFileCategory(),
	}
}

func changeCategoryToComposeFile(args []string) bool {
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
