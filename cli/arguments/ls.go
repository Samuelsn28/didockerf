package arguments

import (
	"fmt"
)

func GetArgLs() Argument {
	return Argument{
		Id: "ls",
		Action: listFiles,
		ValidArgs: make(map[ArgumentId]Argument),
	}

}

func listFiles(args []string) bool {
	fmt.Println("Listando tudo...")
	return true
}

