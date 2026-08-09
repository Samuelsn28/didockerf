package arguments

const dockerfileCategoryID ArgumentID = "dfile"

func GetSubCommandDfile() Argument {
	return Argument{
		ID:        dockerfileCategoryID,
		Action:    changeCategoryToDockerfile,
		ValidArgs: getValidArgsForDockerfileCategory(),
	}
}

func changeCategoryToDockerfile(args []string) bool {
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
