package arguments

import (
	"fmt"
)

func GetSubCommandExport() Argument {
	return Argument{
		Id: "export",
		Action: exportAllFiles,
		ValidArgs: make(map[ArgumentId]Argument),
	}
}

func exportAllFiles([]string) bool {
	fmt.Println("Exportando...")
	
	return true
}

