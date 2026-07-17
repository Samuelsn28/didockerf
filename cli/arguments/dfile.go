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

	return map[ArgumentId]Argument {
		ls.Id: ls,
	}

}

