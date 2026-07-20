package arguments

import (
	"fmt"
)

const composeFileCategoryId ArgumentId = "composefile"

func GetSubCommandComposeFile() Argument {
	return Argument{
		Id: composeFileCategoryId,
		Action: changeCategoryToComposeFile,
		ValidArgs: getValidArgsForComposeFileCategory(),
	}
}

func changeCategoryToComposeFile(args []string) bool {
	fmt.Println("Now working with compose files")

	return false
}

func getValidArgsForComposeFileCategory() map[ArgumentId]Argument {
	ls := GetArgLsComposeFile()
	save := GetArgSaveComposeFile()
	rm := GetArgRmComposeFile()
	ch := GetArgChangeComposeFile()
	prune := GetArgPruneComposeFile()

	return map[ArgumentId]Argument {
		ls.Id: ls,
		save.Id: save,
		rm.Id: rm,
		ch.Id: ch,
		prune.Id: prune,
	}	

}



