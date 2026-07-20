package arguments

import (
	"fmt"
)

const exportId ArgumentId = "export"

func GetSubCommandExport() Argument {
	return Argument{
		Id: exportId,
		Action: exportAllFiles,
		ValidArgs: make(map[ArgumentId]Argument),
	}
}

func exportAllFiles([]string) bool {
	fmt.Println("Exportando...")
	
	return true
}

