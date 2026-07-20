package arguments

import (
	"fmt"
)

const dockerfileCategoryId ArgumentId = "dfile"

func GetSubCommandDfile() Argument {
	return Argument{
		Id: dockerfileCategoryId,
		Action: changeCategoryToDockerfile,
		ValidArgs: getValidArgsForDockerfileCategory(),
	}
}

func changeCategoryToDockerfile(args []string) bool {
	fmt.Println("DFile is here")
	return false
}

func getValidArgsForDockerfileCategory() map[ArgumentId]Argument {
	ls := GetArgLsDockerfile()
	save := GetArgSaveDockerfile()
	rm := GetArgRmDockerfile()
	ch := GetArgChangeDockerfile()
	prune := GetArgPruneDockerfile()

	return map[ArgumentId]Argument {
		ls.Id: ls,
		save.Id: save,
		rm.Id: rm,
		ch.Id: ch,
		prune.Id: prune,
	}	

}



