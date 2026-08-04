package arguments

import (
	"fmt"
)

const composeFileCategoryID ArgumentID = "composefile"

func GetSubCommandComposeFile() Argument {
	return Argument{
		ID:        composeFileCategoryID,
		Action:    changeCategoryToComposeFile,
		ValidArgs: getValidArgsForComposeFileCategory(),
	}
}

func changeCategoryToComposeFile(args []string) bool {
	fmt.Println("Now working with compose files")

	return false
}

func getValidArgsForComposeFileCategory() map[ArgumentID]Argument {
	ls := GetArgLsComposeFile()
	save := GetArgSaveComposeFile()
	rm := GetArgRmComposeFile()
	ch := GetArgChangeComposeFile()
	prune := GetArgPruneComposeFile()

	return map[ArgumentID]Argument{
		ls.ID:    ls,
		save.ID:  save,
		rm.ID:    rm,
		ch.ID:    ch,
		prune.ID: prune,
	}
}
