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
	rm := GetArgRmDockerfile()
	ch := GetArgChangeDockerfile()
	prune := GetArgPruneDockerfile()

	return map[ArgumentID]Argument{
		ls.ID:    ls,
		save.ID:  save,
		rm.ID:    rm,
		ch.ID:    ch,
		prune.ID: prune,
	}
}
