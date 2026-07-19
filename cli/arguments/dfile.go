package arguments

import (
	"fmt"
)

func GetSubCommandDfile() Argument {
	return Argument{
		Id: "dfile",
		Action: changeCategory,
		ValidArgs: getValidArgs(),
	}
}

func changeCategory(args []string) bool {
	fmt.Println("DFile is here")
	return false
}

func getValidArgs() map[ArgumentId]Argument {
	ls := GetArgLs()
	save := GetArgSaveDockerfile()
	rm := GetArgRmDockerfile()
	ch := GetArgChangeDockerfile()

	return map[ArgumentId]Argument {
		ls.Id: ls,
		save.Id: save,
		rm.Id: rm,
		ch.Id: ch,
	}

}



